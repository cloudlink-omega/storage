package types

var GameFeatureTags = map[string]string{

	// Features
	"achievements": "This game supports earning achievements.",
	"config":       "This game can make changes to your account.",
	"controllers":  "This game supports controllers.",
	"dlc":          "This game supports Downloadable Content.",
	"legacy":       "This game supports legacy CloudLink clients.",
	"matchmaking":  "This game supports matchmaking.",
	"save":         "This game supports cloud save data.",
	"points":       "This game can earn, spend, trade or redeem points.",

	// Age Ratings
	"everyone": "This game is suitable for everyone.",
	"mature":   "This game is only for adult audiences.",
	"older":    "This game is only for teens and older audiences.",

	// Platform support
	"mobile":   "This game is only available on mobile devices.",
	"multidev": "This game can be played on mobile or desktop devices.",

	// Source Code
	"oss":         "The game is open source.",
	"proprietary": "The source code of this game is proprietary.",

	// Development Platforms
	"ontw":      "This game was made using Turbowarp.",
	"onpm":      "This game was made using PenguinMod.",
	"oneq":      "This game was made using E羊icques (SheepTester's Mod).",
	"onscratch": "This game is also available on Scratch.",

	// Advisories
	"violent":    "This game contains or references violent content.",
	"substances": "This game contains or references drugs, alcohol or weapons.",

	// Status
	"review": "This game is undergoing review or awaiting approval by an administrator.",

	// Extra Connectivity
	"call":  "This game supports voice chat.",
	"mail":  "This game can send and receive messages using your account.",
	"vchat": "This game supports proximity voice chat.",
	"vmail": "This game supports sending or receiving voicemail.",
}
