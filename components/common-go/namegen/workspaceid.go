// Copyright (c) 2020 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package namegen

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

// PossibleWorkspaceIDPatterns
// nxpod-protocol/src/util/generate-workspace-id.ts is authoritative over the generation
// ws-proxy/pkg/proxy/workspacerouter.go is authoritative for this regexp
var PossibleWorkspaceIDPatterns = []string{
	"[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
	"[0-9a-z]{2,16}-[0-9a-z]{2,16}-[0-9a-z]{8,11}",
}

var workspaceIDPattern = regexp.MustCompile(getWorkspaceIDPatternStr())

// getWorkspaceIDPatternStr is the expected Workspace ID pattern str
// ^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$|^[0-9a-z]{2,16}-[0-9a-z]{2,16}-[0-9a-z]{8,11}$
func getWorkspaceIDPatternStr() string {
	patterns := []string{}
	for _, p := range PossibleWorkspaceIDPatterns {
		patterns = append(patterns, fmt.Sprintf("^%s$", p))
	}
	return strings.Join(patterns, "|")
}

func GenerateWorkspaceID() (string, error) {
	s1, err := chooseRandomly(colors, 1)
	if err != nil {
		return "", err
	}
	s2, err := chooseRandomly(animals, 1)
	if err != nil {
		return "", err
	}
	s3, err := chooseRandomly(characters, 11)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{s1, s2, s3}, "-"), nil
}

var (
	InvalidWorkspaceID = errors.New("workspace id does not match required format")
)

func ValidateWorkspaceID(id string) error {
	if !workspaceIDPattern.MatchString(id) {
		return fmt.Errorf("id '%s' does not match workspace ID regex '%s': %w", id, workspaceIDPattern.String(), InvalidWorkspaceID)
	}

	return nil
}

func chooseRandomly(options []string, length int) (res string, err error) {
	l := big.NewInt(int64(len(options)))
	for i := 0; i < length; i++ {
		c, err := rand.Int(rand.Reader, l)
		if err != nil {
			return "", err
		}
		res += options[int(c.Int64())]
	}
	return res, nil
}

var characters = strings.Split("abcdefghijklmnopqrstuvwxyz0123456789", "")

var colors = []string{
	"amaranth",
	"amber",
	"amethyst",
	"apricot",
	"aqua",
	"aquamarine",
	"azure",
	"beige",
	"black",
	"blue",
	"blush",
	"bronze",
	"brown",
	"chocolate",
	"coffee",
	"copper",
	"coral",
	"crimson",
	"cyan",
	"emerald",
	"fuchsia",
	"gold",
	"gray",
	"green",
	"harlequin",
	"indigo",
	"ivory",
	"jade",
	"kumquat",
	"lavender",
	"lime",
	"magenta",
	"maroon",
	"moccasin",
	"olive",
	"orange",
	"peach",
	"pink",
	"plum",
	"purple",
	"red",
	"rose",
	"salmon",
	"sapphire",
	"scarlet",
	"silver",
	"tan",
	"teal",
	"tomato",
	"turquoise",
	"violet",
	"white",
	"yellow",
}

var animals = []string{
	"canidae",
	"felidae",
	"cat",
	"cattle",
	"dog",
	"donkey",
	"goat",
	"horse",
	"pig",
	"rabbit",
	"aardvark",
	"aardwolf",
	"albatross",
	"alligator",
	"alpaca",
	"amphibian",
	"anaconda",
	"angelfish",
	"anglerfish",
	"ant",
	"anteater",
	"antelope",
	"antlion",
	"ape",
	"aphid",
	"armadillo",
	"asp",
	"baboon",
	"badger",
	"bandicoot",
	"barnacle",
	"barracuda",
	"basilisk",
	"bass",
	"bat",
	"bear",
	"beaver",
	"bedbug",
	"bee",
	"beetle",
	"bird",
	"bison",
	"blackbird",
	"boa",
	"boar",
	"bobcat",
	"bobolink",
	"bonobo",
	"booby",
	"bovid",
	"bug",
	"butterfly",
	"buzzard",
	"camel",
	"canid",
	"capybara",
	"cardinal",
	"caribou",
	"carp",
	"cat",
	"catshark",
	"caterpillar",
	"catfish",
	"cattle",
	"centipede",
	"cephalopod",
	"chameleon",
	"cheetah",
	"chickadee",
	"chicken",
	"chimpanzee",
	"chinchilla",
	"chipmunk",
	"clam",
	"clownfish",
	"cobra",
	"cockroach",
	"cod",
	"condor",
	"constrictor",
	"coral",
	"cougar",
	"cow",
	"coyote",
	"crab",
	"crane",
	"crawdad",
	"crayfish",
	"cricket",
	"crocodile",
	"crow",
	"cuckoo",
	"cicada",
	"damselfly",
	"deer",
	"dingo",
	"dinosaur",
	"dodo",
	"dog",
	"dolphin",
	"donkey",
	"dormouse",
	"dove",
	"dragonfly",
	"dragon",
	"duck",
	"eagle",
	"earthworm",
	"earwig",
	"echidna",
	"eel",
	"egret",
	"elephant",
	"elk",
	"emu",
	"ermine",
	"falcon",
	"ferret",
	"finch",
	"firefly",
	"fish",
	"flamingo",
	"flea",
	"fly",
	"flyingfish",
	"fowl",
	"fox",
	"frog",
	"gamefowl",
	"galliform",
	"gazelle",
	"gecko",
	"gerbil",
	"gibbon",
	"giraffe",
	"goat",
	"goldfish",
	"goose",
	"gopher",
	"gorilla",
	"grasshopper",
	"grouse",
	"guan",
	"guanaco",
	"guineafowl",
	"gull",
	"guppy",
	"haddock",
	"halibut",
	"hamster",
	"hare",
	"harrier",
	"hawk",
	"hedgehog",
	"heron",
	"herring",
	"hippopotamus",
	"hookworm",
	"hornet",
	"horse",
	"hoverfly",
	"hummingbird",
	"hyena",
	"iguana",
	"impala",
	"jackal",
	"jaguar",
	"jay",
	"jellyfish",
	"junglefowl",
	"kangaroo",
	"kingfisher",
	"kite",
	"kiwi",
	"koala",
	"koi",
	"krill",
	"ladybug",
	"lamprey",
	"landfowl",
	"lark",
	"leech",
	"lemming",
	"lemur",
	"leopard",
	"leopon",
	"limpet",
	"lion",
	"lizard",
	"llama",
	"lobster",
	"locust",
	"loon",
	"louse",
	"lungfish",
	"lynx",
	"macaw",
	"mackerel",
	"magpie",
	"mammal",
	"manatee",
	"mandrill",
	"marlin",
	"marmoset",
	"marmot",
	"marsupial",
	"marten",
	"mastodon",
	"meadowlark",
	"meerkat",
	"mink",
	"minnow",
	"mite",
	"mockingbird",
	"mole",
	"mollusk",
	"mongoose",
	"monkey",
	"moose",
	"mosquito",
	"moth",
	"mouse",
	"mule",
	"muskox",
	"narwhal",
	"newt",
	"nightingale",
	"ocelot",
	"octopus",
	"opossum",
	"orangutan",
	"orca",
	"ostrich",
	"otter",
	"owl",
	"ox",
	"panda",
	"panther",
	"parakeet",
	"parrot",
	"parrotfish",
	"partridge",
	"peacock",
	"peafowl",
	"pelican",
	"penguin",
	"perch",
	"pheasant",
	"pig",
	"pigeon",
	"pike",
	"pinniped",
	"piranha",
	"planarian",
	"platypus",
	"pony",
	"porcupine",
	"porpoise",
	"possum",
	"prawn",
	"primate",
	"ptarmigan",
	"puffin",
	"puma",
	"python",
	"quail",
	"quelea",
	"quokka",
	"rabbit",
	"raccoon",
	"rat",
	"rattlesnake",
	"raven",
	"reindeer",
	"reptile",
	"rhinoceros",
	"roadrunner",
	"rodent",
	"rook",
	"rooster",
	"roundworm",
	"sailfish",
	"salamander",
	"salmon",
	"sawfish",
	"scallop",
	"scorpion",
	"seahorse",
	"shark",
	"sheep",
	"shrew",
	"shrimp",
	"silkworm",
	"silverfish",
	"skink",
	"skunk",
	"sloth",
	"slug",
	"smelt",
	"snail",
	"snake",
	"snipe",
	"sole",
	"sparrow",
	"spider",
	"spoonbill",
	"squid",
	"squirrel",
	"starfish",
	"stingray",
	"stoat",
	"stork",
	"sturgeon",
	"swallow",
	"swan",
	"swift",
	"swordfish",
	"swordtail",
	"tahr",
	"takin",
	"tapir",
	"tarantula",
	"tarsier",
	"termite",
	"tern",
	"thrush",
	"tick",
	"tiger",
	"tiglon",
	"toad",
	"tortoise",
	"toucan",
	"trout",
	"tuna",
	"turkey",
	"turtle",
	"tyrannosaurus",
	"urial",
	"vicuna",
	"viper",
	"vole",
	"vulture",
	"wallaby",
	"walrus",
	"wasp",
	"warbler",
	"weasel",
	"whale",
	"whippet",
	"whitefish",
	"wildcat",
	"wildebeest",
	"wildfowl",
	"wolf",
	"wolverine",
	"wombat",
	"woodpecker",
	"worm",
	"wren",
	"xerinae",
	"yak",
	"zebra",
	"alpaca",
	"cat",
	"cattle",
	"chicken",
	"dog",
	"donkey",
	"ferret",
	"gayal",
	"goldfish",
	"guppy",
	"horse",
	"koi",
	"llama",
	"sheep",
	"yak",
	"unicorn",
}
