package vm

import (
	"testing"

	"github.com/end-r/guardian"
)

func TestStorageArrayDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.New(a).CompileString(
		`contract ArrayTest {
            animals = [string]{
                "Dog", "Cat"
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}

func TestStorageMapDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.New(a).CompileString(
		`contract ArrayTest {
            animals = map[string]string{
                "Dog":"canine", "Cat":"feline",
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}

func TestMemoryArrayDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.New(a).CompileString(
		`contract ArrayTest {

            func doThings(){
                animals = [string]{
                    "Dog", "Cat"
                }
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}

func TestMemoryMapDeclaration(t *testing.T) {
	a := new(Arsonist)
	guardian.New(a).CompileString(
		`contract ArrayTest {

            func doThings(){
                animals = map[string]string{
                    "Dog":"canine", "Cat":"feline",
                }
            }
        }
    `)
	checkMnemonics(t, a.VM.Instructions, []string{
		"",
	})
}
