package owasp

// Entry represents a single OWASP LLM Top 10 2025 entry
type Entry struct {
	ID          string
	Name        string
	Description string
	ATLASIDs    []string
}

// Top10 is the OWASP LLM Top 10 for 2025
var Top10 = []Entry{
	{
		ID:   "LLM01",
		Name: "Prompt Injection",
		Description: "Attackers manipulate LLM inputs to override instructions, bypass safety controls, " +
			"or cause unintended actions. Includes direct injection via user input and indirect injection " +
			"via external content the model ingests.",
		ATLASIDs: []string{"AML.T0051", "AML.T0051.000", "AML.T0051.001", "AML.T0051.002"},
	},
	{
		ID:   "LLM02",
		Name: "Sensitive Information Disclosure",
		Description: "LLMs may expose sensitive data through responses, including training data, " +
			"system prompts, PII, credentials, or internal configuration details.",
		ATLASIDs: []string{"AML.T0057", "AML.T0069"},
	},
	{
		ID:   "LLM03",
		Name: "Supply Chain Vulnerabilities",
		Description: "Risks introduced through third-party datasets, pre-trained models, plugins, " +
			"or infrastructure components that may be compromised, outdated, or malicious.",
		ATLASIDs: []string{"AML.T0010", "AML.T0010.000", "AML.T0010.001", "AML.T0010.002", "AML.T0010.003", "AML.T0010.004", "AML.T0010.005"},
	},
	{
		ID:   "LLM04",
		Name: "Data and Model Poisoning",
		Description: "Adversaries corrupt training data or fine-tuning pipelines to embed backdoors, " +
			"bias outputs, or degrade model performance in targeted ways.",
		ATLASIDs: []string{"AML.T0020", "AML.T0043", "AML.T0018"},
	},
	{
		ID:   "LLM05",
		Name: "Improper Output Handling",
		Description: "LLM-generated content is passed to downstream systems without validation. " +
			"Can lead to XSS, SSRF, command injection, or privilege escalation depending on context.",
		ATLASIDs: []string{"AML.T0048", "AML.T0051"},
	},
	{
		ID:   "LLM06",
		Name: "Excessive Agency",
		Description: "LLMs granted excessive permissions or autonomy take unintended actions with " +
			"real-world consequences, especially in agentic pipelines with tool access.",
		ATLASIDs: []string{"AML.T0053", "AML.T0080", "AML.T0099"},
	},
	{
		ID:   "LLM07",
		Name: "System Prompt Leakage",
		Description: "Hidden system instructions are exposed through model responses, revealing " +
			"internal configuration, security controls, or sensitive operational details.",
		ATLASIDs: []string{"AML.T0057", "AML.T0056"},
	},
	{
		ID:   "LLM08",
		Name: "Vector and Embedding Weaknesses",
		Description: "Vulnerabilities in RAG pipelines and vector stores allow adversaries to " +
			"poison retrieval context, extract embeddings, or manipulate semantic search results.",
		ATLASIDs: []string{"AML.T0070", "AML.T0064", "AML.T0066"},
	},
	{
		ID:   "LLM09",
		Name: "Misinformation",
		Description: "LLMs generate plausible but false information, enabling disinformation campaigns, " +
			"fraudulent content, or manipulation of decision-making processes.",
		ATLASIDs: []string{"AML.T0048", "AML.T0019"},
	},
	{
		ID:   "LLM10",
		Name: "Unbounded Consumption",
		Description: "Attackers exploit LLM resource usage through excessive requests, denial of " +
			"service, or prompt-based attacks that trigger expensive model operations.",
		ATLASIDs: []string{"AML.T0034", "AML.T0029"},
	},
}

// Get returns an entry by ID (e.g. "LLM01") or nil if not found
func Get(id string) *Entry {
	for i := range Top10 {
		if Top10[i].ID == id {
			return &Top10[i]
		}
	}
	return nil
}