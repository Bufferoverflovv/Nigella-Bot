package extras

import (
	"log"
	"strconv"
)

const (
	DEFAULT             = 0        // #000000
	AQUA                = 1752220  // #1ABC9C
	DARK_AQUA           = 1146986  // #11806A
	GREEN               = 3066993  // #2ECC71
	DARK_GREEN          = 2067276  // #1F8B4C
	BLUE                = 3447003  // #3498DB
	DARK_BLUE           = 2123412  // #206694
	PURPLE              = 10181046 // #9B59B6
	DARK_PURPLE         = 7419530  // #71368A
	LUMINOUS_VIVID_PINK = 15277667 // #E91E63
	DARK_VIVID_PINK     = 11342935 // #AD1457
	GOLD                = 15844367 //#F1C40F
	DARK_GOLD           = 12745742 // #C27C0E
	ORANGE              = 15105570 // #E67E22
	DARK_ORANGE         = 11027200 // #A84300
	RED                 = 15158332 // #E74C3C
	DARK_RED            = 10038562 // #992D22
	GREY                = 9807270  //#95A5A6
	DARK_GREY           = 9936031  // #979C9F
	DARKER_GREY         = 8359053  //#7F8C8D
	LIGHT_GREY          = 12370112 //#BCC0C0
	NAVY                = 3426654  // #34495E
	DARK_NAVY           = 2899536  //#2C3E50
	YELLOW              = 16776960 // #FFFF00

	// Official Discord Colours
	DISCORD_WHITE      = 16777215 // #FFFFFF
	BLURPLE            = 5793266  // #5865F2
	GREYPLE            = 10070709 // #99AAB5
	DARK_BUT_NOT_BLACK = 2895667  // #2C2F33
	NOT_QUITE_BLACK    = 2303786  //#23272A
	DISCORD_GREEN      = 5763719  // #57F287
	DISCORD_YELLOW     = 16705372 // #FEE75C
	FUSCHIA            = 15418782 // #EB459E
	DISCORD_RED        = 15548997 // #ED4245
	WHITE              = 16777215 // #FFFFFF
	BLACK              = 2303786  // #23272A
)

// Use to convert hex values to decimal for embed colours
func ConvertHex(HexValue string) int {

	num, err := strconv.ParseInt(HexValue, 16, 64)

	if err != nil {
		log.Printf("Not a valid value %s", err)
		return 0
	}

	return int(num)
}
