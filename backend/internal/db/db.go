package db

import (
	"log"
	"os"

	"github.com/nedpals/supabase-go"
)

var Supabase *supabase.Client

// InitDatabase initializes the Supabase client
func InitDatabase() {
	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	if url == "" || key == "" {
		log.Fatal("SUPABASE_URL or SUPABASE_KEY not set in .env")
	}

	Supabase = supabase.CreateClient(url, key)
	log.Println("✅ Connected to Supabase")
}
