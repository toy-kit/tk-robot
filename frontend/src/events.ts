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
        { label: "ReadAll", icon: "iconfont icon-input" },
        { label: "WriteAll", icon: "iconfont icon-input", component: "Str" },
    ],
    Mouse: [
        { label: "Click", icon: "iconfont icon-click", component: "Click" },
        { label: "DoubleClick", icon: "iconfont icon-click", component: "Click" },
        { label: "Move", icon: "iconfont icon-move", component: "Move" },
        { label: "MoveSmooth", icon: "iconfont icon-move", component: "MoveSmooth" },
        { label: "DragSmooth", icon: "iconfont icon-move", component: "DragSmooth" },
        { label: "ScrollDir", icon: "iconfont icon-move", component: "ScrollDir" },
        { label: "Location", icon: "iconfont icon-coord" },
    ],
    Keyboard: [
        { label: "KeyTap", icon: "iconfont icon-input", component: "Keys" },
        { label: "KeyPress", icon: "iconfont icon-input", component: "Str" },
        { label: "KeyDown", icon: "iconfont icon-input", component: "Keys" },
        { label: "KeyUp", icon: "iconfont icon-input", component: "Keys" },
        { label: "TypeStr", icon: "iconfont icon-input", component: "Str" },
        { label: "CopyStr", icon: "iconfont icon-input" },
        { label: "PasteStr", icon: "iconfont icon-input" },
        { label: "CtrlTap", icon: "iconfont icon-input", component: "Keys" },
    ],
    Screen: [
        { label: "CaptureScreen", icon: "iconfont icon-input" },
        { label: "CaptureImg", icon: "iconfont icon-input", component: "AppName" },
        { label: "BitmapFindStr", icon: "iconfont icon-input", component: "Str" },
        { label: "BitmapFindPic", icon: "iconfont icon-input", component: "Path" },
        { label: "BitmapFindColor", icon: "iconfont icon-input", component: "Color" },
    ],
    Window: [
        { label: "Run", icon: "iconfont icon-active", component: "Path" },
        { label: "GetTitle", icon: "iconfont icon-title" },
        { label: "ActiveName", icon: "iconfont icon-active", component: "AppName" },
        { label: "MinWindow", icon: "iconfont icon-active", component: "AppName" },
        { label: "MaxWindow", icon: "iconfont icon-active", component: "AppName" },
        { label: "CloseWindow", icon: "iconfont icon-active", component: "AppName" },
    ]
}