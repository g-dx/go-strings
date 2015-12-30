package main
import (
	"bufio"
	"fmt"
	"os"
	"github.com/g-dx/go-strings/trie"
)

func main() {

	// Create tries
	tries := make(map[string]trie.Trie)
	tries["Simple trie"] = trie.Simple()
	tries["LCRS trie"] = trie.LCRS()
	tries["Radix trie"] = trie.Radix()

	// Load dictionary
	f, err := os.Open("./dict/words.txt")
	if err != nil {
		panic(err)
	}
	stat, err := f.Stat()
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Load dictionary into tries
	n := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for _, t := range tries {
			t.Insert(scanner.Text())
		}
		n++
	}

	// Check for error
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// Print stats
	fmt.Printf("Dictionary Size    : %vKb\n", stat.Size() / 1024)
	fmt.Printf("Dictionary Entries : %v words\n", n)
	fmt.Println()
	fmt.Printf("| %-15v | %-10v | %-15v\n", "Type", "Nodes", "Mem Size (Kb)")
	fmt.Println("=============================================")
	for name, trie := range tries {
		fmt.Printf("| %-15v | %-10v | %-7v\n", name, trie.Count(), trie.SizeOf()/1024)
	}
}