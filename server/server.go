package server

import (
    // ... imports ...
)

// StartAgent est la seule fonction que l'utilisateur du framework appelle.
func StartAgent(agent Agent) {
    // 1. Récupère la carte de l'agent
    card := agent.GetAgentCard()
    log.Printf("Démarrage de l'agent : %s (ID: %s)", card.Name, card.Id)

    // 2. Met en place le serveur ConnectRPC
    mux := http.NewServeMux()
    server := connect.NewServer(mux) // Hypothetical server creation
    
    // 3. L'agent enregistre ses propres services RPC
    agent.RegisterServices(server)

    // 4. LOGIQUE D'ENREGISTREMENT ACTIF (gérée par le framework !)
    go func() {
        registerWithDirectory(card) // Appelle l'annuaire
        startHeartbeat(card.Id)     // Commence à envoyer des signaux de vie
    }()

    // 5. Appelle la logique de démarrage de l'agent
    if err := agent.Startup(); err != nil {
        log.Fatalf("Erreur au démarrage de l'agent : %v", err)
    }

    // 6. Démarre le serveur HTTP et écoute les requêtes
    log.Printf("Agent démarré sur %s", card.Url)
    http.ListenAndServe(card.Url, mux)
}