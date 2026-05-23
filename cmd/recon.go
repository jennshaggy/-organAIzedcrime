package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reconCmd = &cobra.Command{
	Use:   "recon",
	Short: "AI infrastructure reconnaissance reference",
}

var reconPortsCmd = &cobra.Command{
	Use:   "ports",
	Short: "AI-specific ports and services",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== AI Infrastructure Ports ===")
		fmt.Println()
		fmt.Println("--- Model Serving ---")
		fmt.Println("  5000   MLflow Tracking Server")
		fmt.Println("  8000   Triton Inference Server (HTTP) / vLLM / Ollama / Ray")
		fmt.Println("  8001   Triton Inference Server (gRPC)")
		fmt.Println("  8002   Triton Prometheus Metrics")
		fmt.Println("  8080   TorchServe Inference")
		fmt.Println("  8081   TorchServe Management API")
		fmt.Println("  8082   TorchServe Metrics")
		fmt.Println("  8265   Ray Dashboard")
		fmt.Println("  8500   TensorFlow Serving (gRPC)")
		fmt.Println("  8501   TensorFlow Serving (HTTP)")
		fmt.Println("  8888   Jupyter Notebook")
		fmt.Println("  11434  Ollama")
		fmt.Println()
		fmt.Println("--- Vector Databases ---")
		fmt.Println("  6333   Qdrant (HTTP)")
		fmt.Println("  6334   Qdrant (gRPC)")
		fmt.Println("  9000   MinIO / Weaviate")
		fmt.Println()
		fmt.Println("--- Recommended Nmap ---")
		fmt.Println("  nmap -p 5000,6333,6334,8000,8001,8002,8080,8081,8082,8265,8500,8501,8888,9000,11434 -sV <target>")
		fmt.Println()
		fmt.Println("[+] None of these should be internet-facing without explicit intent. Happy hunting.")
	},
}

var reconFingerprintCmd = &cobra.Command{
	Use:   "fingerprint",
	Short: "AI service fingerprinting techniques",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== AI Service Fingerprinting ===")
		fmt.Println()
		fmt.Println("--- Header Fingerprinting ---")
		fmt.Println("  TorchServe     Server: TorchServe/0.x.x")
		fmt.Println("  Triton         NV-Status header present")
		fmt.Println("                 Send 'endpoint-load-metrics-format: text' to get GPU telemetry in headers")
		fmt.Println("  FastAPI/ML     server: uvicorn + routes like /predict or /embeddings")
		fmt.Println("  OpenAI-compat  x-request-id header + JSON with 'object': 'model'")
		fmt.Println()
		fmt.Println("--- API Response Signatures ---")
		fmt.Println("  TensorFlow     {\"model_version_status\": [{\"version\": \"1\", \"state\": \"AVAILABLE\"}]}")
		fmt.Println("  Triton         {\"name\": \"model\", \"versions\": [\"1\"], \"platform\": \"tensorflow_graphdef\"}")
		fmt.Println("  MLflow         Stack traces referencing mlflow.server or mlflow.tracking")
		fmt.Println("  OpenAI-compat  {\"object\": \"model\", \"id\": \"...\", \"created\": ...}")
		fmt.Println()
		fmt.Println("--- Error Message Fingerprinting ---")
		fmt.Println("  TF Serving     Send malformed payload -> error mentions 'tensorinfo_map'")
		fmt.Println("  MLflow         Stack trace references mlflow.server / databricks namespaces")
		fmt.Println("  MLflow CVE     CVE-2024-1558: path traversal exposes server filesystem paths")
		fmt.Println("  Databricks     Java exception io.jsonwebtoken.IncorrectClaimException")
		fmt.Println()
		fmt.Println("--- gRPC Fingerprinting ---")
		fmt.Println("  grpcurl -plaintext <target>:8001 list")
		fmt.Println("  grpcurl -plaintext <target>:8001 describe inference.GRPCInferenceService")
		fmt.Println()
		fmt.Println("[+] Debug-friendly defaults rarely get turned off before production. Exploit that.")
	},
}

var reconEndpointsCmd = &cobra.Command{
	Use:   "endpoints",
	Short: "AI-specific endpoint naming conventions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("=== AI Endpoint Conventions ===")
		fmt.Println()
		fmt.Println("--- Inference ---")
		fmt.Println("  /predict")
		fmt.Println("  /invocations        (SageMaker convention)")
		fmt.Println("  /infer")
		fmt.Println("  /generate")
		fmt.Println("  /embeddings")
		fmt.Println("  /score")
		fmt.Println()
		fmt.Println("--- Model Management ---")
		fmt.Println("  /v1/models")
		fmt.Println("  /v2/models")
		fmt.Println("  /v2/models/<name>/config")
		fmt.Println("  /v2/models/<name>/infer")
		fmt.Println()
		fmt.Println("--- MLflow ---")
		fmt.Println("  /api/2.0/mlflow/registered-models/list")
		fmt.Println("  /api/2.0/mlflow/model-versions/search")
		fmt.Println("  /api/2.0/mlflow/experiments/search")
		fmt.Println("  /api/2.0/mlflow/runs/log-batch")
		fmt.Println()
		fmt.Println("--- Orchestration ---")
		fmt.Println("  /pipeline/apis/v1beta1/     (Kubeflow)")
		fmt.Println("  /api/kernels                (Jupyter)")
		fmt.Println("  /api/contents/              (Jupyter)")
		fmt.Println("  /metrics                    (Prometheus)")
		fmt.Println()
		fmt.Println("--- ATLAS Technique ---")
		fmt.Println("  atlas technique get AML.T0069   Discover LLM System Information")
		fmt.Println("  atlas technique get AML.T0014   Discover ML Model Family")
		fmt.Println()
		fmt.Println("[+] Add these to your ffuf/feroxbuster wordlists. SecLists won't have them.")
	},
}

func init() {
	reconCmd.AddCommand(reconPortsCmd)
	reconCmd.AddCommand(reconFingerprintCmd)
	reconCmd.AddCommand(reconEndpointsCmd)
	rootCmd.AddCommand(reconCmd)
}