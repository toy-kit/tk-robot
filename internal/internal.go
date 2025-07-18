package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/bitmap"
)

type HandlerFunc func(args Params, config Config) (any, error)

var handlers = map[string]HandlerFunc{
	// Base
	"Sleep":               Sleep,
	"LocationByMouseDown": LocationByMouseDown,
	"ColorByMouseDown":    ColorByMouseDown,
	"KeyByListen":         KeyByListen,
	// Mouse
	"Click":       Click,
	"DoubleClick": DoubleClick,
	"Move":        Move,
	"MoveSmooth":  MoveSmooth,
	"DragSmooth":  DragSmooth,
	"ScrollDir":   ScrollDir,
	"Location":    Location,
	// Keyboard
	"KeyTap":   KeyTap,
	"KeyPress": KeyPress,
	"KeyDown":  KeyDown,
	"KeyUp":    KeyUp,
	"ReadAll":  ReadAll,
	"WriteAll": WriteAll,
	"TypeStr":  TypeStr,
	"PasteStr": PasteStr,
	// Screen
	"CaptureScreen":   CaptureScreen,
	"CaptureImg":      CaptureImg,
	"BitmapFindStr":   BitmapFindStr,
	"BitmapFindPic":   BitmapFindPic,
	"BitmapFindColor": BitmapFindColor,
	// Window
	"Run":         Run,
	"ActiveName":  ActiveName,
	"GetTitle":    GetTitle,
	"MinWindow":   MinWindow,
	"MaxWindow":   MaxWindow,
	"CloseWindow": CloseWindow,
}

func Call(name string, args Params, config Config) Result {
	robotgo.MouseSleep = 300
	robotgo.KeySleep = 100
	r := Result{Status: false}
	fun, ok := handlers[name]
	if !ok {
		r.Message = fmt.Sprintf("no such handler: %v", name)
		return r
	}
	d, err := fun(args, config)
	if err != nil {
		r.Message = err.Error()
		return r
	}
	r.Data = d
	r.Status = true
	return r
}

func GetConfig() map[string]any {
	config := map[string]any{}
	w, h := robotgo.GetScreenSize()
	config["w"] = w
	config["h"] = h
	return config
}

// Process get the all process struct
func Process() []map[string]any {
	items := []map[string]any{}
	ps, err := robotgo.Process()
	if err != nil {
		return items
	}
	for _, p := range ps {
		if p.Name != "" && p.Pid != 0 {
			items = append(items, map[string]any{"name": p.Name, "pid": p.Pid})
		}
	}
	return items
}

// ColorByMouseDown get the mouse location color
func ColorByMouseDown(args Params, config Config) (any, error) {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		hook.End()
	})
	s := hook.Start()
	<-hook.Process(s)
	return robotgo.GetLocationColor(), nil
}

// CaptureScreen capture the screen and return image.Image, error
func CaptureScreen(args Params, config Config) (any, error) {
	if err := PathExists("screenshot"); err != nil {
		return nil, err
	}
	w, h := robotgo.GetScreenSize()
	img, err := robotgo.CaptureImg(0, 0, w, h)
	if err != nil {
		return nil, err
	}
	filename := fmt.Sprintf("screenshot/%v.png", time.Now().UnixNano())
	err = robotgo.Save(img, filename)
	return filename, err
}

// CaptureImg capture the screen and return image.Image, error
func CaptureImg(args Params, config Config) (any, error) {
	if err := PathExists("screenshot"); err != nil {
		return nil, err
	}
	if err := checkPid(args.Pid); err != nil {
		return nil, err
	}
	x, y, w, h := robotgo.GetBounds(args.Pid)
	img, err := robotgo.CaptureImg(x, y, w, h)
	if err != nil {
		return nil, err
	}
	filename := fmt.Sprintf("screenshot/%v.png", time.Now().UnixNano())
	err = robotgo.Save(img, filename)
	return filename, err
}
func getScreenBit() robotgo.CBitmap {
	w, h := robotgo.GetScreenSize()
	return robotgo.CaptureScreen(0, 0, w, h)
}

// BitmapFindPic finding the image by path
func BitmapFindPic(args Params, config Config) (any, error) {
	bit := getScreenBit()
	defer robotgo.FreeBitmap(bit)
	x, y := bitmap.FindPic(args.Path, bit)
	return map[string]int{"x": x, "y": y}, nil
}

// BitmapFindColor find bitmap color
func BitmapFindColor(args Params, config Config) (any, error) {
	bit := getScreenBit()
	defer robotgo.FreeBitmap(bit)
	r, g, b, err := color16ToRGB(args.Color)
	if err != nil {
		return nil, err
	}
	x, y := bitmap.FindColor(robotgo.CHex(robotgo.RgbToHex(uint8(r), uint8(g), uint8(b))), bit)
	return map[string]int{"x": x, "y": y}, nil
}

// BitmapFindStr bitmap from string
func BitmapFindStr(args Params, config Config) (any, error) {
	bit := getScreenBit()
	defer robotgo.FreeBitmap(bit)
	bit2 := bitmap.FromStr(args.Str)
	defer robotgo.FreeBitmap(bit2)
	x, y := bitmap.Find(bit2, bit)
	return map[string]int{"x": x, "y": y}, nil
}

// Run run a cmd shell
func Run(args Params, config Config) (any, error) {
	cmdName := "/bin/bash"
	params := "-c"
	if runtime.GOOS == "windows" {
		cmdName = "cmd"
		params = "/c"
	}
	cmd := exec.Command(cmdName, params, args.Path)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	go cmd.Wait()
	robotgo.Sleep(1)
	return map[string]int{"pid": cmd.Process.Pid}, err
}

// ActiveName active the window
func ActiveName(args Params, config Config) (any, error) {
	return nil, checkPid(args.Pid)
}

// MinWindow set the window min
func MinWindow(args Params, config Config) (any, error) {
	if err := checkPid(args.Pid); err != nil {
		return nil, err
	}
	robotgo.MinWindow(args.Pid)
	return nil, nil
}

// MaxWindow set the window max
func MaxWindow(args Params, config Config) (any, error) {
	if err := checkPid(args.Pid); err != nil {
		return nil, err
	}
	robotgo.MaxWindow(args.Pid)
	return nil, nil
}

// CloseWindow close the window
func CloseWindow(args Params, config Config) (any, error) {
	if err := checkPid(args.Pid); err != nil {
		return nil, err
	}
	robotgo.CloseWindow(args.Pid)
	return nil, nil
}

// Move move the mouse to (x, y)
func Move(args Params, config Config) (any, error) {
	robotgo.Move(args.X, args.Y)
	return nil, nil
}

// MoveSmooth move the mouse smooth,
func MoveSmooth(args Params, config Config) (any, error) {
	robotgo.MoveSmooth(args.X, args.Y, 0.1*float64(args.Speed), 0.3*float64(args.Speed))
	return nil, nil
}

// DragSmooth drag the mouse like smooth to (x, y)
func DragSmooth(args Params, config Config) (any, error) {
	robotgo.DragSmooth(args.X, args.Y, 0.1*float64(args.Speed), 0.3*float64(args.Speed))
	return nil, nil
}

// ScrollDir scroll the mouse with direction to (x, "up")
func ScrollDir(args Params, config Config) (any, error) {
	robotgo.ScrollDir(args.X, args.Direction)
	return nil, nil
}

// KeyByListen get the mouse location color
func KeyByListen(args Params, config Config) (any, error) {
	var code uint16
	hook.Register(hook.KeyUp, []string{}, func(e hook.Event) {
		code = e.Keycode
		hook.End()
	})
	s := hook.Start()
	<-hook.Process(s)
	key := ""
	for k, v := range hook.Keycode {
		if v == code {
			key = k
			break
		}
	}
	return key, nil
}

// KeyTap taps the keyboard code;
func KeyTap(args Params, config Config) (any, error) {
	err := robotgo.KeyTap(args.Keys[0], args.Keys[0:])
	return nil, err
}

// KeyPress press key string
func KeyPress(args Params, config Config) (any, error) {
	err := robotgo.KeyPress(args.Str)
	return nil, err
}

// KeyDown press down a key
func KeyDown(args Params, config Config) (any, error) {
	err := robotgo.KeyDown(args.Keys[0])
	return nil, err
}

// KeyUp press up a key
func KeyUp(args Params, config Config) (any, error) {
	err := robotgo.KeyUp(args.Keys[0])
	return nil, err
}

// ReadAll read string from clipboard
func ReadAll(args Params, config Config) (any, error) {
	return robotgo.ReadAll()
}

// WriteAll write string to clipboard
func WriteAll(args Params, config Config) (any, error) {
	err := robotgo.WriteAll(args.Str)
	return nil, err
}

// TypeStr send a string (supported UTF-8)
func TypeStr(args Params, config Config) (any, error) {
	robotgo.TypeStrDelay(args.Str, 100)
	return nil, nil
}

// PasteStr paste a string (support UTF-8),
func PasteStr(args Params, config Config) (any, error) {
	err := robotgo.PasteStr(args.Str)
	return nil, err
}

// Click click the mouse button
func Click(args Params, config Config) (any, error) {
	robotgo.Click(args.Button, false)
	return nil, nil
}

// DoubleClick double-click the mouse button
func DoubleClick(args Params, config Config) (any, error) {
	robotgo.Click(args.Button, true)
	return nil, nil
}

// LocationByMouseDown get the mouse location
func LocationByMouseDown(args Params, config Config) (any, error) {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		hook.End()
	})
	s := hook.Start()
	<-hook.Process(s)
	return Location(args, config)
}

// Location get the mouse location position return x, y
func Location(args Params, config Config) (any, error) {
	x, y := robotgo.Location()
	return map[string]int{"x": x, "y": y}, nil
}

// Sleep time.Sleep tm second
func Sleep(args Params, config Config) (any, error) {
	robotgo.Sleep(args.Tm)
	return nil, nil
}

// GetTitle get the window title return string
func GetTitle(args Params, config Config) (any, error) {
	title := robotgo.GetTitle()
	return title, nil
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if entry == "" {
			continue
		}
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func checkPid(pid int) error {
	robotgo.ActivePid(pid)
	robotgo.MilliSleep(100)
	p := robotgo.GetPid()
	if pid == p {
		return nil
	}
	return fmt.Errorf("pid not found or the window min")
}

func color16ToRGB(colorStr string) (red, green, blue int, err error) {
	color64, err := strconv.ParseInt(strings.TrimPrefix(colorStr, "#"), 16, 32)
	if err != nil {
		return
	}
	colorInt := int(color64)
	return colorInt >> 16, (colorInt & 0x00FF00) >> 8, colorInt & 0x0000FF, nil
}

func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}
