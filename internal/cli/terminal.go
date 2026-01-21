package cli

import (
	"bufio"
	"disver/internal/host"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func StartTerminal(cnf *string, peer *host.Host) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Commands: '/ping [addr] to send ping message'. config: %s\n", *cnf)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "/ping") {
			addr := strings.TrimPrefix(text, "/ping ")
			peer.SendPINGMessage(addr)

		} else if strings.HasPrefix(text, "/iam") {
			fmt.Printf("Node ID anda adalah:\n%s\n", hex.EncodeToString(peer.Node.ID[:]))
		}
	}
}
