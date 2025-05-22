package types

// {key} => {description, isUser, isDeveloper, isGame}
var ReportTags = map[string][]any{

	// All
	"tos":     {"Violates the terms of service.", true, true, true},
	"slander": {"Slanderous or false information.", true, true, true},
	"spam":    {"Contains spam.", true, true, true},
	"ip":      {"Violates my/someone's intellectual property rights.", true, true, true},

	// Game
	"missing":    {"The project file is missing or damaged.", false, false, true},
	"nsfw":       {"Contains NSFW content and is not correctly tagged.", false, false, true},
	"violent":    {"Contains violent content and is not correctly tagged.", false, false, true},
	"substances": {"Contains or refers to substances and is not correctly tagged.", false, false, true},

	// Developer or user
	"bullying": {"Bullying or harassment.", true, true, false},
}
