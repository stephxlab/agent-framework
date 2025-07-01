package agent

import pb "github.com/stephxlab/agent-framework/gen/framework/v1"

// Tout agent qui veut utiliser notre framework DOIT implémenter cette interface.
type Agent interface {
    // GetAgentCard doit retourner la carte d'identité de l'agent.
    GetAgentCard() *pb.AgentCard

    // Services à enregistrer auprès du serveur ConnectRPC.
    // Le développeur listera ici les services RPC que son agent expose.
    RegisterServices(server *connect.Server)

    // Logique de démarrage spécifique à l'agent (si nécessaire).
    Startup() error

    // Logique d'arrêt propre (si nécessaire).
    Shutdown() error
}