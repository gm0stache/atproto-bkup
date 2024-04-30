package testhelper

import (
	"fmt"
	"net"
	"os"
)

const (
	envIPFSNodeHost string = "TEST_IPFS_HOST"
	envIPFSNodePort string = "TEST_IPFS_PORT"
)

// GetIPFSNodePath returns the IPFS path of the node used for testing.
func GetIPFSNodePath() (string, error) {
	port := os.Getenv(envIPFSNodePort)
	if port == "" {
		return "", fmt.Errorf("env var %q must be set", envIPFSNodePort)
	}

	host := os.Getenv(envIPFSNodeHost)
	if host == "" {
		return "", fmt.Errorf("env var %q must be set", envIPFSNodeHost)
	}

	addr := net.JoinHostPort(host, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}

	hostIP, hostPort, err := net.SplitHostPort(conn.RemoteAddr().String())
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("/ip4/%s/tcp/%s", hostIP, hostPort)
	return path, nil
}
