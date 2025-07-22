import Sleep from "./components/Sleep.vue"
import Move from "./components/Move.vue"
import MoveSmooth from "./components/MoveSmooth.vue"
import DragSmooth from "./components/DragSmooth.vue"
import ScrollDir from "./components/ScrollDir.vue"
import AppName from "./components/AppName.vue"
import Click from "./components/Click.vue"
import Str from "./components/Str.vue"
import Path from "./components/Path.vue"
import Color from "./components/Color.vue"
import Keys from "./components/Keys.vue"
export const components: any = {
    "AppName": AppName,
    "Sleep": Sleep,
    "Move": Move,
    "MoveSmooth": MoveSmooth,
    "DragSmooth": DragSmooth,
    "ScrollDir": ScrollDir,
    "GetTitle": Sleep,
    "Location": Sleep,
    "Click": Click,
    "Str": Str,
    "Path": Path,
    "Color": Color,
    "Keys": Keys,
}
export const events: any = {
    Base: [
        { label: "Sleep", icon: "iconfont icon-sleep", component: "Sleep" },
        { label: "ReadAll", icon: "iconfont icon-read_clipboard" },
        { label: "WriteAll", icon: "iconfont icon-write_clipboard", component: "Str" },
    ],
    Mouse: [
        { label: "Click", icon: "iconfont icon-click", component: "Click" },
        { label: "DoubleClick", icon: "iconfont icon-doubleclick", component: "Click" },
        { label: "Move", icon: "iconfont icon-move", component: "Move" },
        { label: "MoveSmooth", icon: "iconfont icon-smoot", component: "MoveSmooth" },
        { label: "DragSmooth", icon: "iconfont icon-drag", component: "DragSmooth" },
        { label: "ScrollDir", icon: "iconfont icon-scroll", component: "ScrollDir" },
        { label: "Location", icon: "iconfont icon-coord" },
    ],
    Keyboard: [
        { label: "KeyTap", icon: "iconfont icon-tap", component: "Keys" },
        { label: "KeyPress", icon: "iconfont icon-keyboard", component: "Str" },
        { label: "KeyDown", icon: "iconfont icon-down", component: "Keys" },
        { label: "KeyUp", icon: "iconfont icon-up", component: "Keys" },
        { label: "TypeStr", icon: "iconfont icon-text", component: "Str" },
        { label: "CopyStr", icon: "iconfont icon-copy" },
        { label: "PasteStr", icon: "iconfont icon-paste" },
        { label: "CtrlTap", icon: "iconfont icon-keyboard", component: "Keys" },
    ],
    Screen: [
        { label: "CaptureScreen", icon: "iconfont icon-printscreen" },
        { label: "CaptureImg", icon: "iconfont icon-screenshot", component: "AppName" },
        { label: "BitmapFindStr", icon: "iconfont icon-target", component: "Str" },
        { label: "BitmapFindPic", icon: "iconfont icon-target", component: "Path" },
        { label: "BitmapFindColor", icon: "iconfont icon-target", component: "Color" },
    ],
    Window: [
        { label: "Run", icon: "iconfont icon-run", component: "Path" },
        { label: "GetTitle", icon: "iconfont icon-title" },
        { label: "ActiveName", icon: "iconfont icon-active", component: "AppName" },
        { label: "MinWindow", icon: "iconfont icon-minimize", component: "AppName" },
        { label: "MaxWindow", icon: "iconfont icon-maximize", component: "AppName" },
        { label: "CloseWindow", icon: "iconfont icon-close", component: "AppName" },
    ]
}