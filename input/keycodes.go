/*  Copyright 2014, Raphael Estrada
    Author email:  <galaktor@gmx.de>
    Project home:  <https://github.com/galaktor/gostwriter>
    Licensed under The GPL v3 License (see README and LICENSE files) */
package input

/*
  #include <linux/input.h>
*/
import "C"

type KeyCode C.__u16

const (
	KEY_RESERVED         = KeyCode(C.KEY_RESERVED)         /* 0     */
	KEY_ESC              = KeyCode(C.KEY_ESC)              /* 1     */
	KEY_1                = KeyCode(C.KEY_1)                /* 2     */
	KEY_2                = KeyCode(C.KEY_2)                /* 3     */
	KEY_3                = KeyCode(C.KEY_3)                /* 4     */
	KEY_4                = KeyCode(C.KEY_4)                /* 5     */
	KEY_5                = KeyCode(C.KEY_5)                /* 6     */
	KEY_6                = KeyCode(C.KEY_6)                /* 7     */
	KEY_7                = KeyCode(C.KEY_7)                /* 8     */
	KEY_8                = KeyCode(C.KEY_8)                /* 9     */
	KEY_9                = KeyCode(C.KEY_9)                /* 10    */
	KEY_0                = KeyCode(C.KEY_0)                /* 11    */
	KEY_MINUS            = KeyCode(C.KEY_MINUS)            /* 12    */
	KEY_EQUAL            = KeyCode(C.KEY_EQUAL)            /* 13    */
	KEY_BACKSPACE        = KeyCode(C.KEY_BACKSPACE)        /* 14    */
	KEY_TAB              = KeyCode(C.KEY_TAB)              /* 15    */
	KEY_Q                = KeyCode(C.KEY_Q)                /* 16    */
	KEY_W                = KeyCode(C.KEY_W)                /* 17    */
	KEY_E                = KeyCode(C.KEY_E)                /* 18    */
	KEY_R                = KeyCode(C.KEY_R)                /* 19    */
	KEY_T                = KeyCode(C.KEY_T)                /* 20    */
	KEY_Y                = KeyCode(C.KEY_Y)                /* 21    */
	KEY_U                = KeyCode(C.KEY_U)                /* 22    */
	KEY_I                = KeyCode(C.KEY_I)                /* 23    */
	KEY_O                = KeyCode(C.KEY_O)                /* 24    */
	KEY_P                = KeyCode(C.KEY_P)                /* 25    */
	KEY_LEFTBRACE        = KeyCode(C.KEY_LEFTBRACE)        /* 26    */
	KEY_RIGHTBRACE       = KeyCode(C.KEY_RIGHTBRACE)       /* 27    */
	KEY_ENTER            = KeyCode(C.KEY_ENTER)            /* 28    */
	KEY_LEFTCTRL         = KeyCode(C.KEY_LEFTCTRL)         /* 29    */
	KEY_A                = KeyCode(C.KEY_A)                /* 30    */
	KEY_S                = KeyCode(C.KEY_S)                /* 31    */
	KEY_D                = KeyCode(C.KEY_D)                /* 32    */
	KEY_F                = KeyCode(C.KEY_F)                /* 33    */
	KEY_G                = KeyCode(C.KEY_G)                /* 34    */
	KEY_H                = KeyCode(C.KEY_H)                /* 35    */
	KEY_J                = KeyCode(C.KEY_J)                /* 36    */
	KEY_K                = KeyCode(C.KEY_K)                /* 37    */
	KEY_L                = KeyCode(C.KEY_L)                /* 38    */
	KEY_SEMICOLON        = KeyCode(C.KEY_SEMICOLON)        /* 39    */
	KEY_APOSTROPHE       = KeyCode(C.KEY_APOSTROPHE)       /* 40    */
	KEY_GRAVE            = KeyCode(C.KEY_GRAVE)            /* 41    */
	KEY_LEFTSHIFT        = KeyCode(C.KEY_LEFTSHIFT)        /* 42    */
	KEY_BACKSLASH        = KeyCode(C.KEY_BACKSLASH)        /* 43    */
	KEY_Z                = KeyCode(C.KEY_Z)                /* 44    */
	KEY_X                = KeyCode(C.KEY_X)                /* 45    */
	KEY_C                = KeyCode(C.KEY_C)                /* 46    */
	KEY_V                = KeyCode(C.KEY_V)                /* 47    */
	KEY_B                = KeyCode(C.KEY_B)                /* 48    */
	KEY_N                = KeyCode(C.KEY_N)                /* 49    */
	KEY_M                = KeyCode(C.KEY_M)                /* 50    */
	KEY_COMMA            = KeyCode(C.KEY_COMMA)            /* 51    */
	KEY_DOT              = KeyCode(C.KEY_DOT)              /* 52    */
	KEY_SLASH            = KeyCode(C.KEY_SLASH)            /* 53    */
	KEY_RIGHTSHIFT       = KeyCode(C.KEY_RIGHTSHIFT)       /* 54    */
	KEY_KPASTERISK       = KeyCode(C.KEY_KPASTERISK)       /* 55    */
	KEY_LEFTALT          = KeyCode(C.KEY_LEFTALT)          /* 56    */
	KEY_SPACE            = KeyCode(C.KEY_SPACE)            /* 57    */
	KEY_CAPSLOCK         = KeyCode(C.KEY_CAPSLOCK)         /* 58    */
	KEY_F1               = KeyCode(C.KEY_F1)               /* 59    */
	KEY_F2               = KeyCode(C.KEY_F2)               /* 60    */
	KEY_F3               = KeyCode(C.KEY_F3)               /* 61    */
	KEY_F4               = KeyCode(C.KEY_F4)               /* 62    */
	KEY_F5               = KeyCode(C.KEY_F5)               /* 63    */
	KEY_F6               = KeyCode(C.KEY_F6)               /* 64    */
	KEY_F7               = KeyCode(C.KEY_F7)               /* 65    */
	KEY_F8               = KeyCode(C.KEY_F8)               /* 66    */
	KEY_F9               = KeyCode(C.KEY_F9)               /* 67    */
	KEY_F10              = KeyCode(C.KEY_F10)              /* 68    */
	KEY_NUMLOCK          = KeyCode(C.KEY_NUMLOCK)          /* 69    */
	KEY_SCROLLLOCK       = KeyCode(C.KEY_SCROLLLOCK)       /* 70    */
	KEY_KP7              = KeyCode(C.KEY_KP7)              /* 71    */
	KEY_KP8              = KeyCode(C.KEY_KP8)              /* 72    */
	KEY_KP9              = KeyCode(C.KEY_KP9)              /* 73    */
	KEY_KPMINUS          = KeyCode(C.KEY_KPMINUS)          /* 74    */
	KEY_KP4              = KeyCode(C.KEY_KP4)              /* 75    */
	KEY_KP5              = KeyCode(C.KEY_KP5)              /* 76    */
	KEY_KP6              = KeyCode(C.KEY_KP6)              /* 77    */
	KEY_KPPLUS           = KeyCode(C.KEY_KPPLUS)           /* 78    */
	KEY_KP1              = KeyCode(C.KEY_KP1)              /* 79    */
	KEY_KP2              = KeyCode(C.KEY_KP2)              /* 80    */
	KEY_KP3              = KeyCode(C.KEY_KP3)              /* 81    */
	KEY_KP0              = KeyCode(C.KEY_KP0)              /* 82    */
	KEY_KPDOT            = KeyCode(C.KEY_KPDOT)            /* 83    */
	KEY_ZENKAKUHANKAKU   = KeyCode(C.KEY_ZENKAKUHANKAKU)   /* 85    */
	KEY_102ND            = KeyCode(C.KEY_102ND)            /* 86    */
	KEY_F11              = KeyCode(C.KEY_F11)              /* 87    */
	KEY_F12              = KeyCode(C.KEY_F12)              /* 88    */
	KEY_RO               = KeyCode(C.KEY_RO)               /* 89    */
	KEY_KATAKANA         = KeyCode(C.KEY_KATAKANA)         /* 90    */
	KEY_HIRAGANA         = KeyCode(C.KEY_HIRAGANA)         /* 91    */
	KEY_HENKAN           = KeyCode(C.KEY_HENKAN)           /* 92    */
	KEY_KATAKANAHIRAGANA = KeyCode(C.KEY_KATAKANAHIRAGANA) /* 93    */
	KEY_MUHENKAN         = KeyCode(C.KEY_MUHENKAN)         /* 94    */
	KEY_KPJPCOMMA        = KeyCode(C.KEY_KPJPCOMMA)        /* 95    */
	KEY_KPENTER          = KeyCode(C.KEY_KPENTER)          /* 96    */
	KEY_RIGHTCTRL        = KeyCode(C.KEY_RIGHTCTRL)        /* 97    */
	KEY_KPSLASH          = KeyCode(C.KEY_KPSLASH)          /* 98    */
	KEY_SYSRQ            = KeyCode(C.KEY_SYSRQ)            /* 99    */
	KEY_RIGHTALT         = KeyCode(C.KEY_RIGHTALT)         /* 100   */
	KEY_LINEFEED         = KeyCode(C.KEY_LINEFEED)         /* 101   */
	KEY_HOME             = KeyCode(C.KEY_HOME)             /* 102   */
	KEY_UP               = KeyCode(C.KEY_UP)               /* 103   */
	KEY_PAGEUP           = KeyCode(C.KEY_PAGEUP)           /* 104   */
	KEY_LEFT             = KeyCode(C.KEY_LEFT)             /* 105   */
	KEY_RIGHT            = KeyCode(C.KEY_RIGHT)            /* 106   */
	KEY_END              = KeyCode(C.KEY_END)              /* 107   */
	KEY_DOWN             = KeyCode(C.KEY_DOWN)             /* 108   */
	KEY_PAGEDOWN         = KeyCode(C.KEY_PAGEDOWN)         /* 109   */
	KEY_INSERT           = KeyCode(C.KEY_INSERT)           /* 110   */
	KEY_DELETE           = KeyCode(C.KEY_DELETE)           /* 111   */
	KEY_MACRO            = KeyCode(C.KEY_MACRO)            /* 112   */
	KEY_MUTE             = KeyCode(C.KEY_MUTE)             /* 113   */
	KEY_VOLUMEDOWN       = KeyCode(C.KEY_VOLUMEDOWN)       /* 114   */
	KEY_VOLUMEUP         = KeyCode(C.KEY_VOLUMEUP)         /* 115   */
	KEY_POWER            = KeyCode(C.KEY_POWER)            /* 116   */
	KEY_KPEQUAL          = KeyCode(C.KEY_KPEQUAL)          /* 117   */
	KEY_KPPLUSMINUS      = KeyCode(C.KEY_KPPLUSMINUS)      /* 118   */
	KEY_PAUSE            = KeyCode(C.KEY_PAUSE)            /* 119   */
	KEY_SCALE            = KeyCode(C.KEY_SCALE)            /* 120   */
	KEY_KPCOMMA          = KeyCode(C.KEY_KPCOMMA)          /* 121   */
	KEY_HANGEUL          = KeyCode(C.KEY_HANGEUL)          /* 122   */
	KEY_HANGUEL          = KeyCode(C.KEY_HANGUEL)          /* KEY_HANGEUL */
	KEY_HANJA            = KeyCode(C.KEY_HANJA)            /* 123   */
	KEY_YEN              = KeyCode(C.KEY_YEN)              /* 124   */
	KEY_LEFTMETA         = KeyCode(C.KEY_LEFTMETA)         /* 125   */
	KEY_RIGHTMETA        = KeyCode(C.KEY_RIGHTMETA)        /* 126   */
	KEY_COMPOSE          = KeyCode(C.KEY_COMPOSE)          /* 127   */
	KEY_STOP             = KeyCode(C.KEY_STOP)             /* 128   */
	KEY_AGAIN            = KeyCode(C.KEY_AGAIN)            /* 129   */
	KEY_PROPS            = KeyCode(C.KEY_PROPS)            /* 130   */
	KEY_UNDO             = KeyCode(C.KEY_UNDO)             /* 131   */
	KEY_FRONT            = KeyCode(C.KEY_FRONT)            /* 132   */
	KEY_COPY             = KeyCode(C.KEY_COPY)             /* 133   */
	KEY_OPEN             = KeyCode(C.KEY_OPEN)             /* 134   */
	KEY_PASTE            = KeyCode(C.KEY_PASTE)            /* 135   */
	KEY_FIND             = KeyCode(C.KEY_FIND)             /* 136   */
	KEY_CUT              = KeyCode(C.KEY_CUT)              /* 137   */
	KEY_HELP             = KeyCode(C.KEY_HELP)             /* 138   */
	KEY_MENU             = KeyCode(C.KEY_MENU)             /* 139   */
	KEY_CALC             = KeyCode(C.KEY_CALC)             /* 140   */
	KEY_SETUP            = KeyCode(C.KEY_SETUP)            /* 141   */
	KEY_SLEEP            = KeyCode(C.KEY_SLEEP)            /* 142   */
	KEY_WAKEUP           = KeyCode(C.KEY_WAKEUP)           /* 143   */
	KEY_FILE             = KeyCode(C.KEY_FILE)             /* 144   */
	KEY_SENDFILE         = KeyCode(C.KEY_SENDFILE)         /* 145   */
	KEY_DELETEFILE       = KeyCode(C.KEY_DELETEFILE)       /* 146   */
	KEY_XFER             = KeyCode(C.KEY_XFER)             /* 147   */
	KEY_PROG1            = KeyCode(C.KEY_PROG1)            /* 148   */
	KEY_PROG2            = KeyCode(C.KEY_PROG2)            /* 149   */
	KEY_WWW              = KeyCode(C.KEY_WWW)              /* 150   */
	KEY_MSDOS            = KeyCode(C.KEY_MSDOS)            /* 151   */
	KEY_COFFEE           = KeyCode(C.KEY_COFFEE)           /* 152   */
	KEY_SCREENLOCK       = KeyCode(C.KEY_SCREENLOCK)       /* KEY_COFFEE */
	KEY_DIRECTION        = KeyCode(C.KEY_DIRECTION)        /* 153   */
	KEY_CYCLEWINDOWS     = KeyCode(C.KEY_CYCLEWINDOWS)     /* 154   */
	KEY_MAIL             = KeyCode(C.KEY_MAIL)             /* 155   */
	KEY_BOOKMARKS        = KeyCode(C.KEY_BOOKMARKS)        /* 156   */
	KEY_COMPUTER         = KeyCode(C.KEY_COMPUTER)         /* 157   */
	KEY_BACK             = KeyCode(C.KEY_BACK)             /* 158   */
	KEY_FORWARD          = KeyCode(C.KEY_FORWARD)          /* 159   */
	KEY_CLOSECD          = KeyCode(C.KEY_CLOSECD)          /* 160   */
	KEY_EJECTCD          = KeyCode(C.KEY_EJECTCD)          /* 161   */
	KEY_EJECTCLOSECD     = KeyCode(C.KEY_EJECTCLOSECD)     /* 162   */
	KEY_NEXTSONG         = KeyCode(C.KEY_NEXTSONG)         /* 163   */
	KEY_PLAYPAUSE        = KeyCode(C.KEY_PLAYPAUSE)        /* 164   */
	KEY_PREVIOUSSONG     = KeyCode(C.KEY_PREVIOUSSONG)     /* 165   */
	KEY_STOPCD           = KeyCode(C.KEY_STOPCD)           /* 166   */
	KEY_RECORD           = KeyCode(C.KEY_RECORD)           /* 167   */
	KEY_REWIND           = KeyCode(C.KEY_REWIND)           /* 168   */
	KEY_PHONE            = KeyCode(C.KEY_PHONE)            /* 169   */
	KEY_ISO              = KeyCode(C.KEY_ISO)              /* 170   */
	KEY_CONFIG           = KeyCode(C.KEY_CONFIG)           /* 171   */
	KEY_HOMEPAGE         = KeyCode(C.KEY_HOMEPAGE)         /* 172   */
	KEY_REFRESH          = KeyCode(C.KEY_REFRESH)          /* 173   */
	KEY_EXIT             = KeyCode(C.KEY_EXIT)             /* 174   */
	KEY_MOVE             = KeyCode(C.KEY_MOVE)             /* 175   */
	KEY_EDIT             = KeyCode(C.KEY_EDIT)             /* 176   */
	KEY_SCROLLUP         = KeyCode(C.KEY_SCROLLUP)         /* 177   */
	KEY_SCROLLDOWN       = KeyCode(C.KEY_SCROLLDOWN)       /* 178   */
	KEY_KPLEFTPAREN      = KeyCode(C.KEY_KPLEFTPAREN)      /* 179   */
	KEY_KPRIGHTPAREN     = KeyCode(C.KEY_KPRIGHTPAREN)     /* 180   */
	KEY_NEW              = KeyCode(C.KEY_NEW)              /* 181   */
	KEY_REDO             = KeyCode(C.KEY_REDO)             /* 182   */
	KEY_F13              = KeyCode(C.KEY_F13)              /* 183   */
	KEY_F14              = KeyCode(C.KEY_F14)              /* 184   */
	KEY_F15              = KeyCode(C.KEY_F15)              /* 185   */
	KEY_F16              = KeyCode(C.KEY_F16)              /* 186   */
	KEY_F17              = KeyCode(C.KEY_F17)              /* 187   */
	KEY_F18              = KeyCode(C.KEY_F18)              /* 188   */
	KEY_F19              = KeyCode(C.KEY_F19)              /* 189   */
	KEY_F20              = KeyCode(C.KEY_F20)              /* 190   */
	KEY_F21              = KeyCode(C.KEY_F21)              /* 191   */
	KEY_F22              = KeyCode(C.KEY_F22)              /* 192   */
	KEY_F23              = KeyCode(C.KEY_F23)              /* 193   */
	KEY_F24              = KeyCode(C.KEY_F24)              /* 194   */
	KEY_PLAYCD           = KeyCode(C.KEY_PLAYCD)           /* 200   */
	KEY_PAUSECD          = KeyCode(C.KEY_PAUSECD)          /* 201   */
	KEY_PROG3            = KeyCode(C.KEY_PROG3)            /* 202   */
	KEY_PROG4            = KeyCode(C.KEY_PROG4)            /* 203   */
	KEY_DASHBOARD        = KeyCode(C.KEY_DASHBOARD)        /* 204   */
	KEY_SUSPEND          = KeyCode(C.KEY_SUSPEND)          /* 205   */
	KEY_CLOSE            = KeyCode(C.KEY_CLOSE)            /* 206   */
	KEY_PLAY             = KeyCode(C.KEY_PLAY)             /* 207   */
	KEY_FASTFORWARD      = KeyCode(C.KEY_FASTFORWARD)      /* 208   */
	KEY_BASSBOOST        = KeyCode(C.KEY_BASSBOOST)        /* 209   */
	KEY_PRINT            = KeyCode(C.KEY_PRINT)            /* 210   */
	KEY_HP               = KeyCode(C.KEY_HP)               /* 211   */
	KEY_CAMERA           = KeyCode(C.KEY_CAMERA)           /* 212   */
	KEY_SOUND            = KeyCode(C.KEY_SOUND)            /* 213   */
	KEY_QUESTION         = KeyCode(C.KEY_QUESTION)         /* 214   */
	KEY_EMAIL            = KeyCode(C.KEY_EMAIL)            /* 215   */
	KEY_CHAT             = KeyCode(C.KEY_CHAT)             /* 216   */
	KEY_SEARCH           = KeyCode(C.KEY_SEARCH)           /* 217   */
	KEY_CONNECT          = KeyCode(C.KEY_CONNECT)          /* 218   */
	KEY_FINANCE          = KeyCode(C.KEY_FINANCE)          /* 219   */
	KEY_SPORT            = KeyCode(C.KEY_SPORT)            /* 220   */
	KEY_SHOP             = KeyCode(C.KEY_SHOP)             /* 221   */
	KEY_ALTERASE         = KeyCode(C.KEY_ALTERASE)         /* 222   */
	KEY_CANCEL           = KeyCode(C.KEY_CANCEL)           /* 223   */
	KEY_BRIGHTNESSDOWN   = KeyCode(C.KEY_BRIGHTNESSDOWN)   /* 224   */
	KEY_BRIGHTNESSUP     = KeyCode(C.KEY_BRIGHTNESSUP)     /* 225   */
	KEY_MEDIA            = KeyCode(C.KEY_MEDIA)            /* 226   */
	KEY_SWITCHVIDEOMODE  = KeyCode(C.KEY_SWITCHVIDEOMODE)  /* 227   */
	KEY_KBDILLUMTOGGLE   = KeyCode(C.KEY_KBDILLUMTOGGLE)   /* 228   */
	KEY_KBDILLUMDOWN     = KeyCode(C.KEY_KBDILLUMDOWN)     /* 229   */
	KEY_KBDILLUMUP       = KeyCode(C.KEY_KBDILLUMUP)       /* 230   */
	KEY_SEND             = KeyCode(C.KEY_SEND)             /* 231   */
	KEY_REPLY            = KeyCode(C.KEY_REPLY)            /* 232   */
	KEY_FORWARDMAIL      = KeyCode(C.KEY_FORWARDMAIL)      /* 233   */
	KEY_SAVE             = KeyCode(C.KEY_SAVE)             /* 234   */
	KEY_DOCUMENTS        = KeyCode(C.KEY_DOCUMENTS)        /* 235   */
	KEY_BATTERY          = KeyCode(C.KEY_BATTERY)          /* 236   */
	KEY_BLUETOOTH        = KeyCode(C.KEY_BLUETOOTH)        /* 237   */
	KEY_WLAN             = KeyCode(C.KEY_WLAN)             /* 238   */
	KEY_UWB              = KeyCode(C.KEY_UWB)              /* 239   */
	KEY_UNKNOWN          = KeyCode(C.KEY_UNKNOWN)          /* 240   */
	KEY_VIDEO_NEXT       = KeyCode(C.KEY_VIDEO_NEXT)       /* 241   */
	KEY_VIDEO_PREV       = KeyCode(C.KEY_VIDEO_PREV)       /* 242   */
	KEY_BRIGHTNESS_CYCLE = KeyCode(C.KEY_BRIGHTNESS_CYCLE) /* 243   */
	KEY_BRIGHTNESS_ZERO  = KeyCode(C.KEY_BRIGHTNESS_ZERO)  /* 244   */
	KEY_DISPLAY_OFF      = KeyCode(C.KEY_DISPLAY_OFF)      /* 245   */
	KEY_WWAN             = KeyCode(C.KEY_WWAN)             /* 246   */
	KEY_WIMAX            = KeyCode(C.KEY_WIMAX)            /* KEY_WWAN */
	KEY_RFKILL           = KeyCode(C.KEY_RFKILL)           /* 247   */
	KEY_MICMUTE          = KeyCode(C.KEY_MICMUTE)          /* 248   */
	KEY_OK               = KeyCode(C.KEY_OK)               /* 0x160 */
	KEY_SELECT           = KeyCode(C.KEY_SELECT)           /* 0x161 */
	KEY_GOTO             = KeyCode(C.KEY_GOTO)             /* 0x162 */
	KEY_CLEAR            = KeyCode(C.KEY_CLEAR)            /* 0x163 */
	KEY_POWER2           = KeyCode(C.KEY_POWER2)           /* 0x164 */
	KEY_OPTION           = KeyCode(C.KEY_OPTION)           /* 0x165 */
	KEY_INFO             = KeyCode(C.KEY_INFO)             /* 0x166 */
	KEY_TIME             = KeyCode(C.KEY_TIME)             /* 0x167 */
	KEY_VENDOR           = KeyCode(C.KEY_VENDOR)           /* 0x168 */
	KEY_ARCHIVE          = KeyCode(C.KEY_ARCHIVE)          /* 0x169 */
	KEY_PROGRAM          = KeyCode(C.KEY_PROGRAM)          /* 0x16a */
	KEY_CHANNEL          = KeyCode(C.KEY_CHANNEL)          /* 0x16b */
	KEY_FAVORITES        = KeyCode(C.KEY_FAVORITES)        /* 0x16c */
	KEY_EPG              = KeyCode(C.KEY_EPG)              /* 0x16d */
	KEY_PVR              = KeyCode(C.KEY_PVR)              /* 0x16e */
	KEY_MHP              = KeyCode(C.KEY_MHP)              /* 0x16f */
	KEY_LANGUAGE         = KeyCode(C.KEY_LANGUAGE)         /* 0x170 */
	KEY_TITLE            = KeyCode(C.KEY_TITLE)            /* 0x171 */
	KEY_SUBTITLE         = KeyCode(C.KEY_SUBTITLE)         /* 0x172 */
	KEY_ANGLE            = KeyCode(C.KEY_ANGLE)            /* 0x173 */
	KEY_ZOOM             = KeyCode(C.KEY_ZOOM)             /* 0x174 */
	KEY_MODE             = KeyCode(C.KEY_MODE)             /* 0x175 */
	KEY_KEYBOARD         = KeyCode(C.KEY_KEYBOARD)         /* 0x176 */
	KEY_SCREEN           = KeyCode(C.KEY_SCREEN)           /* 0x177 */
	KEY_PC               = KeyCode(C.KEY_PC)               /* 0x178 */
	KEY_TV               = KeyCode(C.KEY_TV)               /* 0x179 */
	KEY_TV2              = KeyCode(C.KEY_TV2)              /* 0x17a */
	KEY_VCR              = KeyCode(C.KEY_VCR)              /* 0x17b */
	KEY_VCR2             = KeyCode(C.KEY_VCR2)             /* 0x17c */
	KEY_SAT              = KeyCode(C.KEY_SAT)              /* 0x17d */
	KEY_SAT2             = KeyCode(C.KEY_SAT2)             /* 0x17e */
	KEY_CD               = KeyCode(C.KEY_CD)               /* 0x17f */
	KEY_TAPE             = KeyCode(C.KEY_TAPE)             /* 0x180 */
	KEY_RADIO            = KeyCode(C.KEY_RADIO)            /* 0x181 */
	KEY_TUNER            = KeyCode(C.KEY_TUNER)            /* 0x182 */
	KEY_PLAYER           = KeyCode(C.KEY_PLAYER)           /* 0x183 */
	KEY_TEXT             = KeyCode(C.KEY_TEXT)             /* 0x184 */
	KEY_DVD              = KeyCode(C.KEY_DVD)              /* 0x185 */
	KEY_AUX              = KeyCode(C.KEY_AUX)              /* 0x186 */
	KEY_MP3              = KeyCode(C.KEY_MP3)              /* 0x187 */
	KEY_AUDIO            = KeyCode(C.KEY_AUDIO)            /* 0x188 */
	KEY_VIDEO            = KeyCode(C.KEY_VIDEO)            /* 0x189 */
	KEY_DIRECTORY        = KeyCode(C.KEY_DIRECTORY)        /* 0x18a */
	KEY_LIST             = KeyCode(C.KEY_LIST)             /* 0x18b */
	KEY_MEMO             = KeyCode(C.KEY_MEMO)             /* 0x18c */
	KEY_CALENDAR         = KeyCode(C.KEY_CALENDAR)         /* 0x18d */
	KEY_RED              = KeyCode(C.KEY_RED)              /* 0x18e */
	KEY_GREEN            = KeyCode(C.KEY_GREEN)            /* 0x18f */
	KEY_YELLOW           = KeyCode(C.KEY_YELLOW)           /* 0x190 */
	KEY_BLUE             = KeyCode(C.KEY_BLUE)             /* 0x191 */
	KEY_CHANNELUP        = KeyCode(C.KEY_CHANNELUP)        /* 0x192 */
	KEY_CHANNELDOWN      = KeyCode(C.KEY_CHANNELDOWN)      /* 0x193 */
	KEY_FIRST            = KeyCode(C.KEY_FIRST)            /* 0x194 */
	KEY_LAST             = KeyCode(C.KEY_LAST)             /* 0x195 */
	KEY_AB               = KeyCode(C.KEY_AB)               /* 0x196 */
	KEY_NEXT             = KeyCode(C.KEY_NEXT)             /* 0x197 */
	KEY_RESTART          = KeyCode(C.KEY_RESTART)          /* 0x198 */
	KEY_SLOW             = KeyCode(C.KEY_SLOW)             /* 0x199 */
	KEY_SHUFFLE          = KeyCode(C.KEY_SHUFFLE)          /* 0x19a */
	KEY_BREAK            = KeyCode(C.KEY_BREAK)            /* 0x19b */
	KEY_PREVIOUS         = KeyCode(C.KEY_PREVIOUS)         /* 0x19c */
	KEY_DIGITS           = KeyCode(C.KEY_DIGITS)           /* 0x19d */
	KEY_TEEN             = KeyCode(C.KEY_TEEN)             /* 0x19e */
	KEY_TWEN             = KeyCode(C.KEY_TWEN)             /* 0x19f */
	KEY_VIDEOPHONE       = KeyCode(C.KEY_VIDEOPHONE)       /* 0x1a0 */
	KEY_GAMES            = KeyCode(C.KEY_GAMES)            /* 0x1a1 */
	KEY_ZOOMIN           = KeyCode(C.KEY_ZOOMIN)           /* 0x1a2 */
	KEY_ZOOMOUT          = KeyCode(C.KEY_ZOOMOUT)          /* 0x1a3 */
	KEY_ZOOMRESET        = KeyCode(C.KEY_ZOOMRESET)        /* 0x1a4 */
	KEY_WORDPROCESSOR    = KeyCode(C.KEY_WORDPROCESSOR)    /* 0x1a5 */
	KEY_EDITOR           = KeyCode(C.KEY_EDITOR)           /* 0x1a6 */
	KEY_SPREADSHEET      = KeyCode(C.KEY_SPREADSHEET)      /* 0x1a7 */
	KEY_GRAPHICSEDITOR   = KeyCode(C.KEY_GRAPHICSEDITOR)   /* 0x1a8 */
	KEY_PRESENTATION     = KeyCode(C.KEY_PRESENTATION)     /* 0x1a9 */
	KEY_DATABASE         = KeyCode(C.KEY_DATABASE)         /* 0x1aa */
	KEY_NEWS             = KeyCode(C.KEY_NEWS)             /* 0x1ab */
	KEY_VOICEMAIL        = KeyCode(C.KEY_VOICEMAIL)        /* 0x1ac */
	KEY_ADDRESSBOOK      = KeyCode(C.KEY_ADDRESSBOOK)      /* 0x1ad */
	KEY_MESSENGER        = KeyCode(C.KEY_MESSENGER)        /* 0x1ae */
	KEY_DISPLAYTOGGLE    = KeyCode(C.KEY_DISPLAYTOGGLE)    /* 0x1af */
	KEY_SPELLCHECK       = KeyCode(C.KEY_SPELLCHECK)       /* 0x1b0 */
	KEY_LOGOFF           = KeyCode(C.KEY_LOGOFF)           /* 0x1b1 */
	KEY_DOLLAR           = KeyCode(C.KEY_DOLLAR)           /* 0x1b2 */
	KEY_EURO             = KeyCode(C.KEY_EURO)             /* 0x1b3 */
	KEY_FRAMEBACK        = KeyCode(C.KEY_FRAMEBACK)        /* 0x1b4 */
	KEY_FRAMEFORWARD     = KeyCode(C.KEY_FRAMEFORWARD)     /* 0x1b5 */
	KEY_CONTEXT_MENU     = KeyCode(C.KEY_CONTEXT_MENU)     /* 0x1b6 */
	KEY_MEDIA_REPEAT     = KeyCode(C.KEY_MEDIA_REPEAT)     /* 0x1b7 */
	KEY_10CHANNELSUP     = KeyCode(C.KEY_10CHANNELSUP)     /* 0x1b8 */
	KEY_10CHANNELSDOWN   = KeyCode(C.KEY_10CHANNELSDOWN)   /* 0x1b9 */
	KEY_IMAGES           = KeyCode(C.KEY_IMAGES)           /* 0x1ba */
	KEY_DEL_EOL          = KeyCode(C.KEY_DEL_EOL)          /* 0x1c0 */
	KEY_DEL_EOS          = KeyCode(C.KEY_DEL_EOS)          /* 0x1c1 */
	KEY_INS_LINE         = KeyCode(C.KEY_INS_LINE)         /* 0x1c2 */
	KEY_DEL_LINE         = KeyCode(C.KEY_DEL_LINE)         /* 0x1c3 */
	KEY_FN               = KeyCode(C.KEY_FN)               /* 0x1d0 */
	KEY_FN_ESC           = KeyCode(C.KEY_FN_ESC)           /* 0x1d1 */
	KEY_FN_F1            = KeyCode(C.KEY_FN_F1)            /* 0x1d2 */
	KEY_FN_F2            = KeyCode(C.KEY_FN_F2)            /* 0x1d3 */
	KEY_FN_F3            = KeyCode(C.KEY_FN_F3)            /* 0x1d4 */
	KEY_FN_F4            = KeyCode(C.KEY_FN_F4)            /* 0x1d5 */
	KEY_FN_F5            = KeyCode(C.KEY_FN_F5)            /* 0x1d6 */
	KEY_FN_F6            = KeyCode(C.KEY_FN_F6)            /* 0x1d7 */
	KEY_FN_F7            = KeyCode(C.KEY_FN_F7)            /* 0x1d8 */
	KEY_FN_F8            = KeyCode(C.KEY_FN_F8)            /* 0x1d9 */
	KEY_FN_F9            = KeyCode(C.KEY_FN_F9)            /* 0x1da */
	KEY_FN_F10           = KeyCode(C.KEY_FN_F10)           /* 0x1db */
	KEY_FN_F11           = KeyCode(C.KEY_FN_F11)           /* 0x1dc */
	KEY_FN_F12           = KeyCode(C.KEY_FN_F12)           /* 0x1dd */
	KEY_FN_1             = KeyCode(C.KEY_FN_1)             /* 0x1de */
	KEY_FN_2             = KeyCode(C.KEY_FN_2)             /* 0x1df */
	KEY_FN_D             = KeyCode(C.KEY_FN_D)             /* 0x1e0 */
	KEY_FN_E             = KeyCode(C.KEY_FN_E)             /* 0x1e1 */
	KEY_FN_F             = KeyCode(C.KEY_FN_F)             /* 0x1e2 */
	KEY_FN_S             = KeyCode(C.KEY_FN_S)             /* 0x1e3 */
	KEY_FN_B             = KeyCode(C.KEY_FN_B)             /* 0x1e4 */
	KEY_BRL_DOT1         = KeyCode(C.KEY_BRL_DOT1)         /* 0x1f1 */
	KEY_BRL_DOT2         = KeyCode(C.KEY_BRL_DOT2)         /* 0x1f2 */
	KEY_BRL_DOT3         = KeyCode(C.KEY_BRL_DOT3)         /* 0x1f3 */
	KEY_BRL_DOT4         = KeyCode(C.KEY_BRL_DOT4)         /* 0x1f4 */
	KEY_BRL_DOT5         = KeyCode(C.KEY_BRL_DOT5)         /* 0x1f5 */
	KEY_BRL_DOT6         = KeyCode(C.KEY_BRL_DOT6)         /* 0x1f6 */
	KEY_BRL_DOT7         = KeyCode(C.KEY_BRL_DOT7)         /* 0x1f7 */
	KEY_BRL_DOT8         = KeyCode(C.KEY_BRL_DOT8)         /* 0x1f8 */
	KEY_BRL_DOT9         = KeyCode(C.KEY_BRL_DOT9)         /* 0x1f9 */
	KEY_BRL_DOT10        = KeyCode(C.KEY_BRL_DOT10)        /* 0x1fa */
	KEY_NUMERIC_0        = KeyCode(C.KEY_NUMERIC_0)        /* 0x200 */
	KEY_NUMERIC_1        = KeyCode(C.KEY_NUMERIC_1)        /* 0x201 */
	KEY_NUMERIC_2        = KeyCode(C.KEY_NUMERIC_2)        /* 0x202 */
	KEY_NUMERIC_3        = KeyCode(C.KEY_NUMERIC_3)        /* 0x203 */
	KEY_NUMERIC_4        = KeyCode(C.KEY_NUMERIC_4)        /* 0x204 */
	KEY_NUMERIC_5        = KeyCode(C.KEY_NUMERIC_5)        /* 0x205 */
	KEY_NUMERIC_6        = KeyCode(C.KEY_NUMERIC_6)        /* 0x206 */
	KEY_NUMERIC_7        = KeyCode(C.KEY_NUMERIC_7)        /* 0x207 */
	KEY_NUMERIC_8        = KeyCode(C.KEY_NUMERIC_8)        /* 0x208 */
	KEY_NUMERIC_9        = KeyCode(C.KEY_NUMERIC_9)        /* 0x209 */
	KEY_NUMERIC_STAR     = KeyCode(C.KEY_NUMERIC_STAR)     /* 0x20a */
	KEY_NUMERIC_POUND    = KeyCode(C.KEY_NUMERIC_POUND)    /* 0x20b */
	KEY_CAMERA_FOCUS     = KeyCode(C.KEY_CAMERA_FOCUS)     /* 0x210 */
	KEY_WPS_BUTTON       = KeyCode(C.KEY_WPS_BUTTON)       /* 0x211 */
	KEY_TOUCHPAD_TOGGLE  = KeyCode(C.KEY_TOUCHPAD_TOGGLE)  /* 0x212 */
	KEY_TOUCHPAD_ON      = KeyCode(C.KEY_TOUCHPAD_ON)      /* 0x213 */
	KEY_TOUCHPAD_OFF     = KeyCode(C.KEY_TOUCHPAD_OFF)     /* 0x214 */
	KEY_CAMERA_ZOOMIN    = KeyCode(C.KEY_CAMERA_ZOOMIN)    /* 0x215 */
	KEY_CAMERA_ZOOMOUT   = KeyCode(C.KEY_CAMERA_ZOOMOUT)   /* 0x216 */
	KEY_CAMERA_UP        = KeyCode(C.KEY_CAMERA_UP)        /* 0x217 */
	KEY_CAMERA_DOWN      = KeyCode(C.KEY_CAMERA_DOWN)      /* 0x218 */
	KEY_CAMERA_LEFT      = KeyCode(C.KEY_CAMERA_LEFT)      /* 0x219 */
	KEY_CAMERA_RIGHT     = KeyCode(C.KEY_CAMERA_RIGHT)     /* 0x21a */
	KEY_ATTENDANT_ON     = KeyCode(C.KEY_ATTENDANT_ON)     /* 0x21b */
	KEY_ATTENDANT_OFF    = KeyCode(C.KEY_ATTENDANT_OFF)    /* 0x21c */
	KEY_ATTENDANT_TOGGLE = KeyCode(C.KEY_ATTENDANT_TOGGLE) /* 0x21d */
	KEY_LIGHTS_TOGGLE    = KeyCode(C.KEY_LIGHTS_TOGGLE)    /* 0x21e */
	KEY_ALS_TOGGLE       = KeyCode(C.KEY_ALS_TOGGLE)       /* 0x230 */
	KEY_MIN_INTERESTING  = KeyCode(C.KEY_MIN_INTERESTING)  /* KEY_MUTE */
	KEY_MAX              = KeyCode(C.KEY_MAX)              /* 0x2ff */
	KEY_CNT              = KeyCode(C.KEY_CNT)              /* (KEY_MAX+1) */

)

var ALL_CODES [KEY_CNT]KeyCode = getAllCodes()

func getAllCodes() [KEY_CNT]KeyCode {
	result := [KEY_CNT]KeyCode{}
	for i := 0; i < int(KEY_CNT); i++ {
		result[i] = KeyCode(i)
	}
	return result
}