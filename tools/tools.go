package tools

// Tool represents a security tool mapped to an ATLAS technique
type Tool struct {
	Name        string
	Description string
	Usage       string
}

// ByTechnique maps ATLAS technique IDs to associated tools
var ByTechnique = map[string][]Tool{
	"AML.T0010.003": {
		{
			Name:        "fickling",
			Description: "Decompile, analyze, and backdoor Python pickle files",
			Usage:       "fickling --check safety -p model.pkl",
		},
		{
			Name:        "modelscan",
			Description: "Scan ML model files for malicious code",
			Usage:       "modelscan -p model.h5",
		},
		{
			Name:        "pickletools",
			Description: "Python stdlib disassembler for pickle streams",
			Usage:       "python3 -m pickletools model.pkl",
		},
	},
	"AML.T0010": {
		{
			Name:        "pip-audit",
			Description: "Scan Python dependencies for known vulnerabilities",
			Usage:       "pip-audit -r requirements.txt",
		},
		{
			Name:        "syft",
			Description: "Generate SBOMs for ML projects",
			Usage:       "syft ./project/ --exclude './venv/**' -o cyclonedx-json",
		},
		{
			Name:        "modelscan",
			Description: "Scan ML model files for malicious code",
			Usage:       "modelscan -p model.h5",
		},
	},
	"AML.T0051": {
		{
			Name:        "promptmap",
			Description: "Automated prompt injection testing for LLM applications",
			Usage:       "promptmap --target http://target/api/chat",
		},
		{
			Name:        "garak",
			Description: "LLM vulnerability scanner — probes for prompt injection, jailbreaks, and more",
			Usage:       "garak --model_type openai --model_name gpt-3.5-turbo",
		},
	},
	"AML.T0054": {
		{
			Name:        "garak",
			Description: "LLM vulnerability scanner with jailbreak probes",
			Usage:       "garak --model_type openai --model_name gpt-3.5-turbo --probes jailbreak",
		},
		{
			Name:        "promptmap",
			Description: "Automated prompt injection and jailbreak testing",
			Usage:       "promptmap --target http://target/api/chat --mode jailbreak",
		},
	},
	"AML.T0069": {
		{
			Name:        "grpcurl",
			Description: "Enumerate gRPC services and methods",
			Usage:       "grpcurl -plaintext <target>:8001 list",
		},
		{
			Name:        "nmap",
			Description: "AI infrastructure port scanning",
			Usage:       "nmap -p 5000,8000-8002,8080-8082,8265,8500,8501,8888,9000,11434 -sV <target>",
		},
		{
			Name:        "ffuf",
			Description: "Fuzz AI-specific endpoint conventions",
			Usage:       "ffuf -u http://target/FUZZ -w ai-endpoints.txt",
		},
	},
	"AML.T0020": {
		{
			Name:        "fickling",
			Description: "Inject malicious payloads into pickle files",
			Usage:       "fickling --inject payload.py model.pkl",
		},
		{
			Name:        "modelscan",
			Description: "Detect poisoned model artifacts",
			Usage:       "modelscan -p candidate_model.pkl",
		},
	},
	"AML.T0057": {
		{
			Name:        "garak",
			Description: "Probe for sensitive data leakage from LLMs",
			Usage:       "garak --probes leakage",
		},
	},
	"AML.T0070": {
		{
			Name:        "curl",
			Description: "Inject into RAG feedback pipelines",
			Usage:       "curl -X POST http://target/api/feedback -H 'Content-Type: application/json' -d '{\"example\":\"[payload]\"}'",
		},
	},
}

// Get returns tools for a given technique ID, checking exact match then parent
func Get(id string) []Tool {
	if t, ok := ByTechnique[id]; ok {
		return t
	}
	return nil
}