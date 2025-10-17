package main

import (
	"fmt"
	"strings"
)

type scenarioContext struct {
	ProgramName      string
	Industry         string
	MethodMix        string
	Constraint       string
	StakeholderGroup string
	Geography        string
	RegFramework     string
	SponsorPersona   string
	TeamProfile      string
	Threat           string
	ValueDriver      string
}

type templateChoice struct {
	Text      string
	Correct   bool
	Rationale string
}

type pmpTemplate struct {
	name        string
	domain      string
	isMulti     bool
	prompt      func(ctx scenarioContext) string
	choices     func(ctx scenarioContext) []templateChoice
	explanation func(ctx scenarioContext, correct []templateChoice) string
	popularity  float64
}

var scenarioContexts = []scenarioContext{
	{
		ProgramName:      "Aurora ERP Unification",
		Industry:         "global manufacturing conglomerate",
		MethodMix:        "hybrid agile with scaled Scrum aligned to Stage-Gate governance",
		Constraint:       "quarterly Sarbanes-Oxley attestations",
		StakeholderGroup: "regional CFO council and unionized plant managers",
		Geography:        "German, Polish, Malaysian, and Brazilian operations",
		RegFramework:     "SOX and EU Digital Operations Act thresholds",
		SponsorPersona:   "data-driven COO who demands weekly benefit proofs",
		TeamProfile:      "outsourced integrator, internal architects, and change leads",
		Threat:           "integration defects delaying consolidated revenue recognition",
		ValueDriver:      "accelerated month-end close and a defensible audit trail",
	},
	{
		ProgramName:      "Helios Digital Claims Platform",
		Industry:         "multinational insurance provider",
		MethodMix:        "scaled Kanban synchronized with predictive release gates",
		Constraint:       "NAIC solvency reporting and GDPR audit windows",
		StakeholderGroup: "regional compliance leads and underwriting directors",
		Geography:        "North America, the UK, and Singapore",
		RegFramework:     "GDPR and NAIC Model Audit Rule obligations",
		SponsorPersona:   "risk-averse Chief Risk Officer demanding airtight controls",
		TeamProfile:      "internal engineers, offshore quality lab, and process excellence coaches",
		Threat:           "legacy policy data failing conversion tests for top-tier clients",
		ValueDriver:      "launch automated straight-through claims decisioning",
	},
	{
		ProgramName:      "Summit Omni-Channel Banking Rollout",
		Industry:         "consumer banking division",
		MethodMix:        "dual-track agile integrated with mandatory waterfall vendor gates",
		Constraint:       "central bank cyber-resilience attestations every four weeks",
		StakeholderGroup: "branch operations leads and cybersecurity architects",
		Geography:        "Middle East and North Africa branches",
		RegFramework:     "Basel III operational risk guidelines",
		SponsorPersona:   "ambitious Chief Digital Officer chasing an aggressive launch date",
		TeamProfile:      "fintech vendor, internal security architects, and change ambassadors",
		Threat:           "penetration-test gaps blocking go-live approval",
		ValueDriver:      "increase digital onboarding by forty percent within twelve months",
	},
	{
		ProgramName:      "Atlas Logistics Control Tower",
		Industry:         "global freight forwarder",
		MethodMix:        "predictive wave planning augmented with agile sprints for analytics",
		Constraint:       "quarterly supply-chain resilience reporting to shipping regulators",
		StakeholderGroup: "regional operations directors and carrier relationship managers",
		Geography:        "strategic hubs in Rotterdam, Dubai, and Shenzhen",
		RegFramework:     "IMO emissions disclosures and customs trade regulations",
		SponsorPersona:   "operations-focused COO who escalates immediately on service impacts",
		TeamProfile:      "system integrator, data scientists, and legacy warehouse SMEs",
		Threat:           "lack of real-time carrier reliability data undermining commitments",
		ValueDriver:      "visibility dashboards that cut diversion costs by twenty percent",
	},
	{
		ProgramName:      "Pulse Telehealth Expansion",
		Industry:         "integrated healthcare network",
		MethodMix:        "agile squads delivering within a predictive clinical release calendar",
		Constraint:       "HIPAA audits and state medical board submissions",
		StakeholderGroup: "chief medical information officers and clinical practice councils",
		Geography:        "United States western region and rural outreach clinics",
		RegFramework:     "HIPAA, ONC interoperability rule, and state telemedicine statutes",
		SponsorPersona:   "physician-CEO focused on patient outcomes and equity",
		TeamProfile:      "EMR vendor, human-centered design team, and clinical super-users",
		Threat:           "clinician adoption stalls due to competing training commitments",
		ValueDriver:      "virtual visit throughput that safeguards chronic-care revenue",
	},
	{
		ProgramName:      "Nova Smart Grid Modernization",
		Industry:         "national energy utility",
		MethodMix:        "hybrid predictive project controls with agile analytics workstreams",
		Constraint:       "grid-resilience mandates from the energy regulator",
		StakeholderGroup: "grid operations chiefs and regulatory affairs directors",
		Geography:        "high-risk coastal and desert substations across three countries",
		RegFramework:     "ISO 55000 asset management and NERC cyber standards",
		SponsorPersona:   "policy-minded sustainability minister scrutinizing benefits",
		TeamProfile:      "engineering contractors, internal data platform team, and field crews",
		Threat:           "IoT sensor failures degrading outage prediction accuracy",
		ValueDriver:      "predictive maintenance insights that defer capital expenditure",
	},
	{
		ProgramName:      "Catalyst Trial Data Platform",
		Industry:         "global biopharma",
		MethodMix:        "scaled agile with validated systems lifecycle checkpoints",
		Constraint:       "FDA and EMA inspection-readiness audits",
		StakeholderGroup: "clinical operations VPs and pharmacovigilance leads",
		Geography:        "phase III sites across the US, EU, and India",
		RegFramework:     "21 CFR Part 11 and ICH GCP guidance",
		SponsorPersona:   "science-driven Chief Medical Officer demanding evidence integrity",
		TeamProfile:      "CRO partner, data managers, and quality assurance auditors",
		Threat:           "inconsistent site data jeopardizing submission timelines",
		ValueDriver:      "regulatory-grade data lineage with faster safety signal detection",
	},
	{
		ProgramName:      "Lighthouse Public ERP Renewal",
		Industry:         "federal civilian agency",
		MethodMix:        "predictive acquisition milestones with agile configuration sprints",
		Constraint:       "appropriations reporting each fiscal quarter",
		StakeholderGroup: "bureau chiefs, unions, and oversight committees",
		Geography:        "multi-state service centers and Washington, D.C. headquarters",
		RegFramework:     "OMB circulars, FITARA scorecards, and GAO audit criteria",
		SponsorPersona:   "scrutinized agency administrator balancing politics and outcomes",
		TeamProfile:      "prime contractor, PMO analysts, and senior civil servants",
		Threat:           "change fatigue and resistance undermining rollout readiness",
		ValueDriver:      "single source financial management and transparent grants tracking",
	},
	{
		ProgramName:      "Vertex Retail Analytics Pivot",
		Industry:         "global specialty retail",
		MethodMix:        "agile product pods feeding a predictive seasonal planning calendar",
		Constraint:       "quarterly earnings guidance and privacy assessments",
		StakeholderGroup: "merchandising VPs and regional e-commerce leaders",
		Geography:        "flagship stores in New York, Paris, Dubai, and Seoul",
		RegFramework:     "CCPA, GDPR, and PCI compliance",
		SponsorPersona:   "results-driven Chief Merchandising Officer tracking conversion lift",
		TeamProfile:      "data scientists, CRM engineers, and omnichannel change partners",
		Threat:           "model drift reducing recommendation accuracy before holiday peak",
		ValueDriver:      "personalized offers that defend margin and loyalty growth",
	},
	{
		ProgramName:      "Orion Defense Mission Suite",
		Industry:         "national defense integrator",
		MethodMix:        "predictive EVMS with agile software increments for avionics",
		Constraint:       "earned value reporting to the defense acquisition executive",
		StakeholderGroup: "program executive office, flight test commanders, and security compliance",
		Geography:        "development centers in Arizona, Ohio, and classified overseas sites",
		RegFramework:     "DFARS cyber requirements and ITAR controls",
		SponsorPersona:   "mission-focused program executive intolerant of slippage",
		TeamProfile:      "prime contractor, avionics engineers, and security accreditation leads",
		Threat:           "requirements churn from rapid threat intelligence updates",
		ValueDriver:      "deployable mission capabilities that meet readiness targets",
	},
	{
		ProgramName:      "Momentum EdTech Personalization",
		Industry:         "global education technology firm",
		MethodMix:        "agile squads orchestrated with predictive academic release windows",
		Constraint:       "accessibility and data privacy audits each semester",
		StakeholderGroup: "curriculum directors and customer success executives",
		Geography:        "North America, LATAM, and ASEAN institutions",
		RegFramework:     "COPPA, FERPA, and regional accessibility mandates",
		SponsorPersona:   "growth-driven Chief Customer Officer promising net retention gain",
		TeamProfile:      "product designers, AI researchers, and implementation coaches",
		Threat:           "pilot institutions threatening churn due to change saturation",
		ValueDriver:      "adaptive learning journeys that improve student outcomes",
	},
	{
		ProgramName:      "Zenith Hospitality CRM",
		Industry:         "luxury hospitality group",
		MethodMix:        "agile CRM releases synchronized with predictive capital projects",
		Constraint:       "brand-standard audits and loyalty program commitments",
		StakeholderGroup: "regional GMs, loyalty marketing, and guest experience labs",
		Geography:        "resorts across the Caribbean, Bali, and the Alps",
		RegFramework:     "GDPR, PCI, and regional tourism authority rules",
		SponsorPersona:   "brand-obsessed Chief Experience Officer tracking NPS",
		TeamProfile:      "CRM vendor, digital concierge team, and on-property champions",
		Threat:           "fragmented guest profiles undermining personalization pilots",
		ValueDriver:      "higher luxury upsell conversion and guest satisfaction",
	},
	{
		ProgramName:      "Equinox Automotive OTA Platform",
		Industry:         "next-gen electric vehicle manufacturer",
		MethodMix:        "scaled agile with predictive homologation checkpoints",
		Constraint:       "transport ministry type-approval windows and cybersecurity rules",
		StakeholderGroup: "vehicle safety regulators and customer service executives",
		Geography:        "assembly plants in Scandinavia, Detroit, and Shenzhen",
		RegFramework:     "UNECE WP.29 software update regulations",
		SponsorPersona:   "innovation-centric CEO promising differentiated driver experience",
		TeamProfile:      "embedded software squads, reliability labs, and customer advocacy teams",
		Threat:           "field data showing critical OTA rollback rates after launch",
		ValueDriver:      "stable OTA pipeline that protects brand trust and revenue streams",
	},
	{
		ProgramName:      "Harbor Port Automation",
		Industry:         "smart port authority",
		MethodMix:        "predictive infrastructure sequencing with agile analytics overlays",
		Constraint:       "harbor-master safety certifications and customs readiness",
		StakeholderGroup: "terminal operators, labor unions, and customs enforcement",
		Geography:        "mega terminals in Singapore, Rotterdam, and Los Angeles",
		RegFramework:     "ILO labor accords and international maritime security codes",
		SponsorPersona:   "publicly visible port director balancing labor and competitiveness",
		TeamProfile:      "automation OEM, data engineers, and operational excellence coaches",
		Threat:           "automation pilots stalling due to labor concerns and training gaps",
		ValueDriver:      "throughput gains while honoring labor agreements and safety",
	},
	{
		ProgramName:      "Quasar Aerospace Avionics Refresh",
		Industry:         "commercial aerospace",
		MethodMix:        "EVMS predictive schedules with agile safety-critical software sprints",
		Constraint:       "FAA and EASA certification checkpoints",
		StakeholderGroup: "flight test directors, supply chain leads, and airline launch customers",
		Geography:        "design center in Seattle with integration labs in Toulouse and Nagoya",
		RegFramework:     "DO-178C and AS9100",
		SponsorPersona:   "engineering-focused program SVP promising no certification slips",
		TeamProfile:      "tier-one suppliers, avionics engineers, and quality compliance",
		Threat:           "supplier firmware defects creating cascading rework",
		ValueDriver:      "certified avionics packages that unlock launch revenue",
	},
}

var pmpTemplates = []pmpTemplate{
	{
		name:       "Program risk and replanning",
		domain:     "People",
		isMulti:    true,
		popularity: 3.9,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("You lead the %s serving a %s across %s. The program blends %s yet must still satisfy %s. Midway through the current increment, %s warn that %s, threatening %s. %s expects you to protect the business case without slipping compliance milestones. Select TWO actions you should take immediately.",
				ctx.ProgramName, ctx.Industry, ctx.Geography, ctx.MethodMix, ctx.Constraint, ctx.StakeholderGroup, ctx.Threat, ctx.ValueDriver, ctx.SponsorPersona)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Facilitate a focused dependency and risk re-planning workshop with %s and leads from %s to refit the roadmap before PI commitments lock.", ctx.StakeholderGroup, ctx.TeamProfile),
					Correct:   true,
					Rationale: "Re-baselining collaboratively keeps stakeholders aligned while protecting flow and transparency.",
				},
				{
					Text:      "Escalate directly to the enterprise steering committee demanding emergency funding before any impact analysis is completed.",
					Correct:   false,
					Rationale: "Escalating without analysis sidesteps core PMI expectations for ownership and data-driven recommendations.",
				},
				{
					Text:      fmt.Sprintf("Quantify the exposure in the risk register, assign mitigation owners from %s, and inject those actions into the program backlog tied to %s.", ctx.TeamProfile, ctx.RegFramework),
					Correct:   true,
					Rationale: "Documenting and integrating mitigations ensures the threat is actively managed and auditable.",
				},
				{
					Text:      "Order product owners to downplay the risk so teams can keep their current velocity metrics untouched.",
					Correct:   false,
					Rationale: "Masking risk violates PMI guidance on transparency and impedes informed decision making.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("PMI expects the project manager to surface significant threats, co-create a response, and embed mitigations into the delivery plan. %s", reasons)
		},
	},
	{
		name:       "Regulatory data remediation",
		domain:     "Process",
		isMulti:    true,
		popularity: 3.85,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("During the %s, %s flag that %s is jeopardizing inspection readiness. You must preserve trust with %s while still meeting %s. Which TWO actions should you lead first?",
				ctx.ProgramName, ctx.StakeholderGroup, ctx.Threat, ctx.SponsorPersona, ctx.Constraint)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Stand up a cross-functional remediation swarm with compliance, data, and delivery leads to isolate root causes and define control measures tied to %s.", ctx.RegFramework),
					Correct:   true,
					Rationale: "Coordinated problem solving clarifies ownership and protects regulatory outcomes.",
				},
				{
					Text:      "Freeze all downstream testing and communicate that go-live will slip until further notice without offering a recovery path.",
					Correct:   false,
					Rationale: "Halting work without plan undermines stakeholder confidence and offers no constructive path forward.",
				},
				{
					Text:      fmt.Sprintf("Update the integrated release plan with the remediation backlog and socialize impacts with %s using scenario-based projections.", ctx.SponsorPersona),
					Correct:   true,
					Rationale: "Transparent replanning with quantified options maintains sponsor confidence and control.",
				},
				{
					Text:      "Delegate the issue entirely to the vendor and notify the regulator that the schedule is at risk before assessing data.",
					Correct:   false,
					Rationale: "Deflecting accountability contradicts PMI leadership expectations and may panic regulators unnecessarily.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("Effective regulatory response requires a structured remediation swarm and transparent replanning with the sponsor. %s", reasons)
		},
	},
	{
		name:       "Vendor performance reset",
		domain:     "Business Environment",
		isMulti:    true,
		popularity: 3.83,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("A core vendor on the %s continues missing integration deliverables, triggering %s concerns. Contract penalties exist, but %s wants measurable recovery before considering replacements. Select TWO responses that align to PMI guidance.",
				ctx.ProgramName, ctx.StakeholderGroup, ctx.SponsorPersona)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Review contract commitments with procurement and the vendor lead, clarifying exit criteria and inserting a joint improvement backlog owned by %s.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Clarifying obligations while building a transparent improvement plan resets expectations without immediate escalation.",
				},
				{
					Text:      "Immediately issue default notices and suspend all vendor work even though no internal contingency exists yet.",
					Correct:   false,
					Rationale: "Punitive action without continuity planning jeopardizes delivery and may breach good-faith obligations.",
				},
				{
					Text:      fmt.Sprintf("Establish weekly performance huddles focused on risk burndown and decision latency, documenting actions in the shared RAID log for %s.", ctx.RegFramework),
					Correct:   true,
					Rationale: "High-frequency transparency ensures issues are owned, tracked, and auditable.",
				},
				{
					Text:      "Silence internal stakeholders about the vendor issues to avoid distracting the sponsor until go-live.",
					Correct:   false,
					Rationale: "Withholding information undermines governance and contradicts PMI expectations for candor.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("Resetting vendor performance means tightening governance and creating a visible recovery backlog before punitive moves. %s", reasons)
		},
	},
	{
		name:       "Distributed team cadence",
		domain:     "People",
		isMulti:    true,
		popularity: 3.82,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Your %s delivery teams span %s. Velocity has dropped because ceremonies clash with mission-critical operations. How do you respond? Select TWO options.", ctx.ProgramName, ctx.Geography)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Re-time key ceremonies so critical time zones rotate convenience, and add asynchronous backlog refinement with decision journals accessible to %s.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Tailored cadence and asynchronous collaboration restore inclusion without losing transparency.",
				},
				{
					Text:      "Mandate that every team attends a single 4 a.m. headquarters stand-up to prove commitment.",
					Correct:   false,
					Rationale: "Ignoring time-zone realities encourages disengagement and attrition.",
				},
				{
					Text:      "Empower local facilitators to run daily syncs feeding a program-level Kanban, and track flow metrics to confirm improvements.",
					Correct:   true,
					Rationale: "Local ownership with flow metrics gives autonomy while preserving oversight.",
				},
				{
					Text:      "Freeze backlog changes until a co-location budget appears, even if work stalls for weeks.",
					Correct:   false,
					Rationale: "Waiting for co-location sacrifices progress and contradicts adaptive leadership.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("PMI-aligned servant leadership adapts cadence and empowers distributed facilitators to keep delivery sustainable. %s", reasons)
		},
	},
	{
		name:       "Knowledge transfer acceleration",
		domain:     "People",
		isMulti:    true,
		popularity: 3.84,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Critical SMEs on the %s will roll off in six weeks; %s fear capability gaps. What TWO steps should you lead?", ctx.ProgramName, ctx.StakeholderGroup)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Stand up a structured knowledge-transfer plan that pairs departing SMEs with successors, linking outputs to the learning backlog owned by %s.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Formal pairing and backlog visibility preserves tacit knowledge.",
				},
				{
					Text:      "Delay current sprint commitments so SMEs can create documentation later when time appears.",
					Correct:   false,
					Rationale: "Deferring knowledge capture compounds risk and misses the window to transfer context.",
				},
				{
					Text:      fmt.Sprintf("Integrate enablement objectives into PI planning and track readiness metrics on the program dashboard for %s.", ctx.SponsorPersona),
					Correct:   true,
					Rationale: "Making enablement visible secures sponsor support and ensures accountability.",
				},
				{
					Text:      "Inform stakeholders that attrition is unavoidable and knowledge loss will be handled after go-live.",
					Correct:   false,
					Rationale: "Accepting loss without mitigation contradicts proactive PMI practices.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("Proactive pairing and transparent readiness tracking retain knowledge before SMEs exit. %s", reasons)
		},
	},
	{
		name:       "Benefits erosion",
		domain:     "Business Environment",
		isMulti:    false,
		popularity: 3.78,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Executive review of the %s shows %s drifting. Forecasts indicate the promised %s could slip by two quarters. What should you do FIRST?", ctx.ProgramName, ctx.ValueDriver, ctx.ValueDriver)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Ask finance to revise benefits targets downward to avoid awkward conversations.",
					Correct:   false,
					Rationale: "Lowering targets prematurely hides issues and ignores root-cause analysis.",
				},
				{
					Text:      fmt.Sprintf("Facilitate a benefits register review with %s to analyze assumptions, update leading indicators, and agree on corrective experiments.", ctx.StakeholderGroup),
					Correct:   true,
					Rationale: "Revalidating benefits with stakeholders aligns the program to value delivery guidance.",
				},
				{
					Text:      "Redirect the team to focus purely on feature throughput and worry about benefits after launch.",
					Correct:   false,
					Rationale: "Separating delivery from outcomes contradicts PMP focus on value.",
				},
				{
					Text:      "Escalate to the board claiming the original business case was flawed to protect yourself.",
					Correct:   false,
					Rationale: "Blame-shifting erodes trust and offers no improvement path.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "PMI guidance centers on continual benefits validation with stakeholders before escalating or adjusting baselines."
		},
	},
	{
		name:       "Sponsor scope challenge",
		domain:     "People",
		isMulti:    false,
		popularity: 3.76,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Halfway through the %s, %s demand a new feature set that jeopardizes the baseline. How do you respond FIRST?", ctx.ProgramName, ctx.SponsorPersona)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Accept the request immediately so the sponsor stays happy, then tell the team later.",
					Correct:   false,
					Rationale: "Agreeing without analysis undermines scope governance.",
				},
				{
					Text:      fmt.Sprintf("Run an impact assessment with %s, evaluate schedule, cost, and benefit trade-offs, and present scenario options through the change control process.", ctx.StakeholderGroup),
					Correct:   true,
					Rationale: "Integrated change control backed by data maintains credibility and sponsor partnership.",
				},
				{
					Text:      "Escalate to the board accusing the sponsor of scope creep before any conversation occurs.",
					Correct:   false,
					Rationale: "Escalating prematurely inflames relationships and skips due diligence.",
				},
				{
					Text:      "Ignore the sponsor's request because the schedule is already stressful.",
					Correct:   false,
					Rationale: "Ignoring stakeholders contradicts PMI expectations for engagement.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "The project manager must analyze impacts and channel the request through formal change control before committing to scope shifts."
		},
	},
	{
		name:       "Risk response gap",
		domain:     "Process",
		isMulti:    false,
		popularity: 3.77,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Risk reviews on the %s show several high threats with no response owners. What should you do?", ctx.ProgramName)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Facilitate a rapid risk-response workshop with %s to assign owners, define triggers, and integrate actions into the plan.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Driving response planning aligns with PMI's expectation for proactive risk management.",
				},
				{
					Text:      "Leave the register as-is; the team already knows the risks exist.",
					Correct:   false,
					Rationale: "Documented responses are required to manage exposure effectively.",
				},
				{
					Text:      "Delete the high risks so the heat map looks better for senior leadership.",
					Correct:   false,
					Rationale: "Hiding risks contradicts transparency and ethics.",
				},
				{
					Text:      "Inform the sponsor that risk management is the PMO's job and not yours.",
					Correct:   false,
					Rationale: "Avoiding accountability violates the project manager's role.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "PMI expects the project manager to ensure high risks carry response strategies with clear ownership and triggers."
		},
	},
	{
		name:       "Metrics misalignment",
		domain:     "Business Environment",
		isMulti:    false,
		popularity: 3.74,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("The %s dashboard shows velocity improving, yet stakeholder satisfaction drops. Which action is BEST aligned with PMP guidance?", ctx.ProgramName)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Revisit KPIs with %s to balance throughput metrics with outcome measures tied to %s.", ctx.StakeholderGroup, ctx.ValueDriver),
					Correct:   true,
					Rationale: "Balanced metrics prevent local optimization and align to value delivery.",
				},
				{
					Text:      "Ignore stakeholder feedback because velocity proves the teams are productive.",
					Correct:   false,
					Rationale: "Productivity without outcomes indicates misaligned success criteria.",
				},
				{
					Text:      "Replace the product owner for lacking positivity.",
					Correct:   false,
					Rationale: "Blame without analysis misses the systemic issue.",
				},
				{
					Text:      "Report the old metrics only so long as the sponsor feels confident.",
					Correct:   false,
					Rationale: "Masking metrics violates transparency.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "Project managers must ensure metrics reflect stakeholder value rather than raw throughput alone."
		},
	},
	{
		name:       "Team conflict mediation",
		domain:     "People",
		isMulti:    false,
		popularity: 3.73,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("On the %s, internal architects and the %s vendor blame each other for integration slippage, and collaboration has stalled. What should you do first?", ctx.ProgramName, ctx.TeamProfile)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Escalate immediately to legal counsel to threaten the vendor with penalties.",
					Correct:   false,
					Rationale: "Escalating to legal before understanding root causes damages relationships.",
				},
				{
					Text:      fmt.Sprintf("Facilitate a joint working session to inspect fact-based data, clarify interfaces, and agree on decision rights aligned to %s expectations.", ctx.RegFramework),
					Correct:   true,
					Rationale: "Resolving conflict with data and shared agreements restores collaboration.",
				},
				{
					Text:      "Assign blame to the internal team to appease the vendor and move on.",
					Correct:   false,
					Rationale: "Assigning blame without evidence erodes trust on both sides.",
				},
				{
					Text:      "Ignore the conflict; velocity will eventually recover on its own.",
					Correct:   false,
					Rationale: "Neglecting conflict allows delays and poor quality to persist.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "Guided facilitation with objective data is required to resolve team conflict and resume joint delivery."
		},
	},
	{
		name:       "Change saturation",
		domain:     "Business Environment",
		isMulti:    false,
		popularity: 3.72,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Regional leaders on the %s warn that %s are exhausted by overlapping initiatives, risking adoption of %s. What should you do?", ctx.ProgramName, ctx.StakeholderGroup, ctx.ValueDriver)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Re-sequence change releases with the change management lead, building a heat map of impacts and coordinating with other portfolios owned by %s.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Coordinated change scheduling protects adoption and reduces fatigue.",
				},
				{
					Text:      "Dismiss the concerns; adoption issues can be fixed after go-live.",
					Correct:   false,
					Rationale: "Ignoring readiness risks undermines value realization.",
				},
				{
					Text:      "Replace local leaders with those more enthusiastic about change.",
					Correct:   false,
					Rationale: "Swapping leaders avoids addressing the systemic overload.",
				},
				{
					Text:      "Delay all other enterprise programs so yours can proceed unchallenged.",
					Correct:   false,
					Rationale: "Demanding priority without negotiation inflames politics.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "Managing change saturation through coordinated planning keeps stakeholders engaged and supports adoption."
		},
	},
	{
		name:       "Quality gating",
		domain:     "Process",
		isMulti:    false,
		popularity: 3.75,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("A quality gate on the %s reveals unresolved defects, yet the sponsor urges keeping the launch date. What's the correct response?", ctx.ProgramName)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Present defect data and risk scenarios to %s, recommending a controlled go/no-go review with remediation options before proceeding.", ctx.SponsorPersona),
					Correct:   true,
					Rationale: "Transparent decision support honors governance while enabling informed trade-offs.",
				},
				{
					Text:      "Ignore the quality gate and proceed so the schedule stays intact.",
					Correct:   false,
					Rationale: "Bypassing gates risks compliance and stakeholder trust.",
				},
				{
					Text:      "Hide the defects from the sponsor to avoid delay debates.",
					Correct:   false,
					Rationale: "Withholding information violates PMI ethics and governance.",
				},
				{
					Text:      "Blame the QA lead publicly so the team works harder before launch.",
					Correct:   false,
					Rationale: "Scapegoating erodes morale and doesn't fix issues.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "Quality gates exist to provide data for informed decisions; the PM must surface facts and facilitate risk-based choices."
		},
	},
	{
		name:       "Procurement alignment",
		domain:     "Business Environment",
		isMulti:    false,
		popularity: 3.71,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Mid-contract, the %s vendor proposes scope changes that increase cost. What's your first move?", ctx.ProgramName)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Approve the change verbally to keep the relationship positive and inform procurement later.",
					Correct:   false,
					Rationale: "Commitments without analysis undermine controls.",
				},
				{
					Text:      fmt.Sprintf("Engage procurement and legal to review the contract, assess change implications, and follow the agreed change mechanism with %s.", ctx.StakeholderGroup),
					Correct:   true,
					Rationale: "Adhering to procurement strategy protects governance and relationships.",
				},
				{
					Text:      "Ignore the vendor's request to teach them a lesson.",
					Correct:   false,
					Rationale: "Stonewalling stalls progress and damages collaboration.",
				},
				{
					Text:      "Tell the team to absorb the extra effort without adjustments.",
					Correct:   false,
					Rationale: "Unplanned scope increases risk burnout and benefits erosion.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "Procurement changes must route through the contract governance process before commitments are made."
		},
	},
	{
		name:       "Release readiness",
		domain:     "Process",
		isMulti:    true,
		popularity: 3.86,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Final release prep for the %s exposes unresolved cross-team defects and missing playbooks. %s still press for launch. Select TWO responses that reflect PMP best practice.", ctx.ProgramName, ctx.SponsorPersona)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Document the issues but move to production so the sponsor sees momentum.",
					Correct:   false,
					Rationale: "Shipping with known gaps without plan risks failure and reputational damage.",
				},
				{
					Text:      fmt.Sprintf("Conduct an integrated readiness review with %s to validate exit criteria, capture defects, and agree go/no-go scenarios.", ctx.StakeholderGroup),
					Correct:   true,
					Rationale: "Collaborative readiness checks ensure stakeholders own launch risk decisions.",
				},
				{
					Text:      fmt.Sprintf("Sequence a short hardening sprint focused on the critical defects, and update run-books with %s before resubmitting to governance.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Targeted closure of high-risk gaps protects launch quality and compliance.",
				},
				{
					Text:      "Delegate the decision entirely to operations without sharing project context.",
					Correct:   false,
					Rationale: "Abdicating context prevents informed operational decisions.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("Launch readiness demands inclusive reviews and targeted remediation prior to committing. %s", reasons)
		},
	},
	{
		name:       "Cost forecast discipline",
		domain:     "Business Environment",
		isMulti:    true,
		popularity: 3.84,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Forecasting on the %s shows CPI dropping while scope remains fixed. Finance wants clarity before approving additional funding. Which TWO actions do you take?", ctx.ProgramName)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Request more budget immediately without updating any forecasts.",
					Correct:   false,
					Rationale: "Asking for funds without data undermines credibility.",
				},
				{
					Text:      fmt.Sprintf("Recalculate EAC scenarios with the finance partner, highlighting trade-offs to protect %s.", ctx.ValueDriver),
					Correct:   true,
					Rationale: "Transparent forecasting options support informed funding decisions.",
				},
				{
					Text:      fmt.Sprintf("Identify cost drivers with %s, create corrective actions, and monitor CPI trend in the dashboard.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Root-cause analysis and monitoring align to proactive cost management.",
				},
				{
					Text:      "Shift contingency from unrelated programs without telling anyone.",
					Correct:   false,
					Rationale: "Moving funds secretly breaches governance and ethics.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("PMI cost management expects data-driven EAC refreshes and targeted corrective actions before seeking additional funding. %s", reasons)
		},
	},
	{
		name:       "Stakeholder communication realignment",
		domain:     "People",
		isMulti:    true,
		popularity: 3.8,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Feedback shows %s receive sporadic updates from the %s, breeding uncertainty about %s. What TWO steps should you take?", ctx.StakeholderGroup, ctx.ProgramName, ctx.ValueDriver)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Rebaseline the stakeholder engagement plan with segmented messaging, cadence, and ownership mapped to %s.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Tailored communication plans keep diverse stakeholders informed and engaged.",
				},
				{
					Text:      "Send a single email assuring everyone that updates will resume later.",
					Correct:   false,
					Rationale: "One-off communication fails to resolve systemic gaps.",
				},
				{
					Text:      fmt.Sprintf("Introduce quarterly value reviews with %s to link progress metrics to %s and capture feedback.", ctx.SponsorPersona, ctx.ValueDriver),
					Correct:   true,
					Rationale: "Interactive reviews reinforce outcomes and create two-way dialogue.",
				},
				{
					Text:      "Limit communications to steering committee members only to avoid noise.",
					Correct:   false,
					Rationale: "Restricting communication fosters misinformation among key stakeholders.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("Sustained stakeholder engagement requires segmented planning and routine value dialogue. %s", reasons)
		},
	},
	{
		name:       "Learning backlog prioritization",
		domain:     "Process",
		isMulti:    true,
		popularity: 3.79,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Retrospectives on the %s reveal recurring spikes in defect leakage due to new tooling. Teams feel pressured to skip improvement work. Select TWO PMI-aligned responses.", ctx.ProgramName)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Negotiate capacity each iteration for improvement stories linked to the learning backlog, and track outcomes on the program dashboard for %s.", ctx.SponsorPersona),
					Correct:   true,
					Rationale: "Allocating capacity institutionalizes continuous improvement.",
				},
				{
					Text:      "Insist teams work overtime indefinitely to handle defects outside normal iterations.",
					Correct:   false,
					Rationale: "Overtime masks systemic issues and risks burnout.",
				},
				{
					Text:      fmt.Sprintf("Pair tool coaches from %s with squads to coach new practices and adjust definition-of-done criteria.", ctx.TeamProfile),
					Correct:   true,
					Rationale: "Building capability and updating quality criteria reduces repeat defects.",
				},
				{
					Text:      "Freeze all feature delivery until the tooling vendor can guarantee zero defects.",
					Correct:   false,
					Rationale: "Halting delivery wholly is unnecessary when improvements can be woven into flow.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			reasons := collectRationales(correct)
			return fmt.Sprintf("Continuous improvement is enabled by planned capacity and coaching support, not overtime heroics. %s", reasons)
		},
	},
	{
		name:       "Regulatory remediation",
		domain:     "Business Environment",
		isMulti:    false,
		popularity: 3.78,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("A surprise audit on the %s uncovers gaps against %s. What is your immediate step?", ctx.ProgramName, ctx.RegFramework)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      fmt.Sprintf("Integrate remediation tasks into the master schedule and risk register, aligning owners from %s and briefing %s on exposure and timelines.", ctx.TeamProfile, ctx.SponsorPersona),
					Correct:   true,
					Rationale: "Embedding remediation into plan and risk processes ensures closure and transparency.",
				},
				{
					Text:      "Dispute the audit findings publicly to protect the team's reputation.",
					Correct:   false,
					Rationale: "Arguing without facts damages credibility.",
				},
				{
					Text:      "Promise auditors the issues will be fixed after go-live without recording them.",
					Correct:   false,
					Rationale: "Verbal promises without traceability risk reoccurrence and sanctions.",
				},
				{
					Text:      "Assign the problem to a single junior analyst to avoid distracting the team.",
					Correct:   false,
					Rationale: "Under-resourcing remediation ignores the severity of audit findings.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "Audit findings require integrated remediation planning with clear ownership and sponsor transparency."
		},
	},
	{
		name:       "Benefits reporting",
		domain:     "Business Environment",
		isMulti:    false,
		popularity: 3.7,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("Steering committee members on the %s request more transparent benefits dashboards tied to %s. What should you do?", ctx.ProgramName, ctx.ValueDriver)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Provide cumulative spend charts only; benefits will emerge post-implementation.",
					Correct:   false,
					Rationale: "Spend without benefits insight limits decision making.",
				},
				{
					Text:      fmt.Sprintf("Co-design a benefits realization dashboard with %s, mapping leading indicators to targets and updating ownership.", ctx.StakeholderGroup),
					Correct:   true,
					Rationale: "Co-created dashboards ensure alignment on outcomes and accountability.",
				},
				{
					Text:      "Wait for finance to produce a report months later.",
					Correct:   false,
					Rationale: "Delaying insights prevents timely steering interventions.",
				},
				{
					Text:      "Share raw operational data without context so the committee can interpret themselves.",
					Correct:   false,
					Rationale: "Uninterpreted data can lead to misinformed conclusions.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "Benefits dashboards must be collaboratively developed to show leading indicators and ownership."
		},
	},
	{
		name:       "Risk escalation",
		domain:     "Process",
		isMulti:    false,
		popularity: 3.76,
		prompt: func(ctx scenarioContext) string {
			return fmt.Sprintf("A critical risk on the %s has triggered despite mitigations and now exceeds your authority threshold. What is the next PMI-aligned action?", ctx.ProgramName)
		},
		choices: func(ctx scenarioContext) []templateChoice {
			return []templateChoice{
				{
					Text:      "Document the breach, escalate to the steering committee with options, and recommend a decision.",
					Correct:   true,
					Rationale: "Escalation with data and options fulfills governance expectations.",
				},
				{
					Text:      "Ignore the threshold; the risk might disappear on its own.",
					Correct:   false,
					Rationale: "Avoidance worsens exposure.",
				},
				{
					Text:      "Transfer the risk to a vendor without contract authority.",
					Correct:   false,
					Rationale: "Unilateral transfer without contractual basis is ineffective.",
				},
				{
					Text:      "Publicly criticize the risk owner for failing to prevent escalation.",
					Correct:   false,
					Rationale: "Blame culture does not resolve the risk.",
				},
			}
		},
		explanation: func(ctx scenarioContext, correct []templateChoice) string {
			return "When authority limits are exceeded, the project manager must escalate with options so governance bodies can decide."
		},
	},
}

var pmpQuestionBank = generatePMPQuestionBank()

func GetPMPQuestions() []QuestionData {
	if len(pmpQuestionBank) <= 200 {
		return append([]QuestionData(nil), pmpQuestionBank...)
	}
	return append([]QuestionData(nil), pmpQuestionBank[:200]...)
}

func GetAdditionalPMPQuestions() []QuestionData {
	if len(pmpQuestionBank) <= 200 {
		return []QuestionData{}
	}
	return append([]QuestionData(nil), pmpQuestionBank[200:]...)
}

func generatePMPQuestionBank() []QuestionData {
	questions := make([]QuestionData, 0, len(pmpTemplates)*len(scenarioContexts))
	for _, tmpl := range pmpTemplates {
		for idx, ctx := range scenarioContexts {
			q := renderTemplate(tmpl, ctx, idx)
			questions = append(questions, q)
		}
	}
	if len(questions) > 300 {
		questions = questions[:300]
	}
	return questions
}

func renderTemplate(tmpl pmpTemplate, ctx scenarioContext, idx int) QuestionData {
	baseChoices := tmpl.choices(ctx)
	labels := []string{"A", "B", "C", "D", "E"}
	choices := make([]ChoiceData, 0, len(baseChoices))
	correct := make([]templateChoice, 0)
	for i, choice := range baseChoices {
		label := labels[i%len(labels)]
		choices = append(choices, ChoiceData{
			text:      choice.Text,
			label:     label,
			isCorrect: choice.Correct,
		})
		if choice.Correct {
			correct = append(correct, choice)
		}
	}

	explanation := tmpl.explanation(ctx, correct)
	if explanation == "" {
		explanation = fmt.Sprintf("Correct responses: %s.", correctOptionSummary(correct))
	} else {
		explanation = fmt.Sprintf("%s Correct responses: %s.", explanation, correctOptionSummary(correct))
	}

	pop := tmpl.popularity + float64(idx%3)*0.01

	return QuestionData{
		prompt:          tmpl.prompt(ctx),
		domain:          tmpl.domain,
		explanation:     explanation,
		popularityScore: pop,
		isMultiSelect:   tmpl.isMulti,
		choices:         choices,
	}
}

func correctOptionSummary(correct []templateChoice) string {
	if len(correct) == 0 {
		return "None"
	}
	snippets := make([]string, 0, len(correct))
	for _, c := range correct {
		snippets = append(snippets, c.Text)
	}
	return strings.Join(snippets, " | ")
}

func collectRationales(correct []templateChoice) string {
	rationales := make([]string, 0, len(correct))
	for _, c := range correct {
		if c.Rationale != "" {
			rationales = append(rationales, c.Rationale)
		}
	}
	if len(rationales) == 0 {
		return ""
	}
	return strings.Join(rationales, " ")
}
