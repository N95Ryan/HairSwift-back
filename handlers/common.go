package handlers

import (
	supa "github.com/supabase-community/supabase-go"
)

var supabaseClient *supa.Client

// InitHandlers initialise les handlers avec le client Supabase
func InitHandlers(client *supa.Client) {
	supabaseClient = client
}
