//Compiler of Golang is stupid need pointer for make them smarter
module not_workout/tests

go 1.22.2

replace not_workout/cmd/permutations => ../cmd/permutations

replace not_workout/cmd/oddInt => ../cmd/oddInt

replace not_workout/cmd/smiley => ../cmd/smiley

require (
	not_workout/cmd/oddInt v0.0.0-00010101000000-000000000000
	not_workout/cmd/permutations v0.0.0-00010101000000-000000000000
	not_workout/cmd/smiley v0.0.0-00010101000000-000000000000
)
