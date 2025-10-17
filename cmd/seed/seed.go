package main

import (
	"fmt"
	"math"
	"strings"
)

// GetQuestions returns all the main CAPM questions based on PMBOK 7th Edition and PMI ECO 2024
func GetQuestions() []QuestionData {
	questions := []QuestionData{
		// Project Management Fundamentals (36%)
		{
			prompt:          "You are managing a project that involves both predictive and adaptive approaches. The stakeholders are requesting frequent changes to requirements. Which project approach would be MOST appropriate for managing the evolving requirements while maintaining project governance?",
			domain:          "Project Management Fundamentals",
			explanation:     "A hybrid approach combines the structure of predictive methods with the flexibility of adaptive approaches, allowing for governance while accommodating changing requirements. This aligns with PMBOK 7th Edition's emphasis on tailoring approaches.",
			popularityScore: 2.8,
			choices: []ChoiceData{
				{"Purely predictive approach with strict change control", "A", false},
				{"Hybrid approach combining predictive and adaptive elements", "B", true},
				{"Fully adaptive approach without any predictive elements", "C", false},
				{"Traditional waterfall with extended planning phase", "D", false},
			},
		},
		{
			prompt:          "According to PMBOK 7th Edition, which statement BEST describes the relationship between project performance domains?",
			domain:          "Project Management Fundamentals",
			explanation:     "PMBOK 7th Edition emphasizes that the eight performance domains are interconnected and interdependent, not independent silos. They work together to deliver project outcomes.",
			popularityScore: 2.9,
			choices: []ChoiceData{
				{"Performance domains operate independently of each other", "A", false},
				{"Each domain has a specific sequence that must be followed", "B", false},
				{"Performance domains are interconnected and interdependent", "C", true},
				{"Only certain domains apply to specific project types", "D", false},
			},
		},
		{
			prompt:          "A project manager is working in a matrix organization where team members report to both functional managers and the project manager. The functional manager is reassigning a key team member to another project. What should the project manager do FIRST?",
			domain:          "Project Management Fundamentals",
			explanation:     "In matrix organizations, negotiation and collaboration with functional managers is essential. The project manager should first discuss the impact and explore alternatives before escalating or making unilateral decisions.",
			popularityScore: 2.7,
			choices: []ChoiceData{
				{"Escalate the issue to senior management immediately", "A", false},
				{"Discuss the impact with the functional manager and negotiate alternatives", "B", true},
				{"Accept the reassignment and find a replacement independently", "C", false},
				{"Refuse to release the team member without approval", "D", false},
			},
		},
		{
			prompt:          "Which of the following BEST represents the concept of 'value' in the context of project outcomes according to PMBOK 7th Edition?",
			domain:          "Project Management Fundamentals",
			explanation:     "PMBOK 7th Edition emphasizes that value is determined by stakeholders and encompasses benefits, worth, and importance beyond just financial measures. Value is stakeholder-dependent and context-specific.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"The financial return on investment of the project", "A", false},
				{"The worth, importance, or benefit of something to stakeholders", "B", true},
				{"The cost savings achieved through efficient project execution", "C", false},
				{"The technical specifications delivered by the project", "D", false},
			},
		},

		// Predictive Plan-Based Methodologies (17%)
		{
			prompt:          "During the planning process group, a project manager discovers that two critical activities can be performed in parallel, but they both require the same specialized resource. The resource is only available 50% of the time. What is the BEST approach to handle this constraint?",
			domain:          "Predictive Methodologies",
			explanation:     "Resource leveling is the technique used to resolve resource conflicts by adjusting the schedule. In this case, the activities may need to be sequential rather than parallel, or additional resources found.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Crash the schedule by adding more resources to both activities", "A", false},
				{"Use resource leveling to adjust the schedule and resolve the conflict", "B", true},
				{"Perform risk analysis and accept the resource conflict", "C", false},
				{"Fast track the activities to reduce the overall duration", "D", false},
			},
		},
		{
			prompt:          "A project has an earned value (EV) of $100,000, actual cost (AC) of $120,000, and planned value (PV) of $90,000. Based on these values, what can you conclude about the project performance?",
			domain:          "Predictive Methodologies",
			explanation:     "CPI = EV/AC = 100,000/120,000 ≈ 0.83, meaning the project is over budget because only 83 cents of value are earned per dollar spent. SPI = EV/PV = 100,000/90,000 ≈ 1.11, so work is ahead of schedule.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"The project is under budget and behind schedule", "A", false},
				{"The project is over budget but ahead of schedule", "B", true},
				{"The project is on budget and on schedule", "C", false},
				{"The project is under budget and ahead of schedule", "D", false},
			},
		},
		{
			prompt:          "During a monthly review, the planned value (PV) is $150,000 and earned value (EV) is $135,000. What is the schedule variance (SV) and what does it indicate?",
			domain:          "Predictive Methodologies",
			explanation:     "SV = EV - PV = 135,000 - 150,000 = -15,000. A negative SV means the project is behind schedule.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"SV = +$15,000; the project is ahead of schedule", "A", false},
				{"SV = -$15,000; the project is behind schedule", "B", true},
				{"SV = -$15,000; the project is under budget", "C", false},
				{"SV = +$15,000; the project is behind schedule", "D", false},
			},
		},
		{
			prompt:          "Midway through execution, EV is reported as $320,000 while PV is $300,000. What is the schedule variance (SV) and how should the project manager describe progress?",
			domain:          "Predictive Methodologies",
			explanation:     "SV = EV - PV = 320,000 - 300,000 = +20,000, indicating the project is ahead of schedule.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"SV = +$20,000; the project is ahead of schedule", "A", true},
				{"SV = +$20,000; the project is over budget", "B", false},
				{"SV = -$20,000; the project is ahead of schedule", "C", false},
				{"SV = -$20,000; the project is behind schedule", "D", false},
			},
		},
		{
			prompt:          "Your project reports an earned value (EV) of $220,000 and an actual cost (AC) of $240,000. What is the cost variance (CV) and how should the performance be interpreted?",
			domain:          "Predictive Methodologies",
			explanation:     "CV = EV - AC = 220,000 - 240,000 = -20,000. A negative CV indicates the project is over budget.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"CV = -$20,000; the project is over budget", "A", true},
				{"CV = +$20,000; the project is under budget", "B", false},
				{"CV = +$20,000; the project is ahead of schedule", "C", false},
				{"CV = -$20,000; the project is behind schedule", "D", false},
			},
		},
		{
			prompt:          "Quarterly reporting shows EV is $510,000 and AC totals $480,000. What is the cost variance (CV) and what does it reveal about cost performance?",
			domain:          "Predictive Methodologies",
			explanation:     "CV = EV - AC = 510,000 - 480,000 = +30,000, meaning the project is under budget.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"CV = +$30,000; the project is under budget", "A", true},
				{"CV = +$30,000; the project is over budget", "B", false},
				{"CV = -$30,000; the project is under budget", "C", false},
				{"CV = -$30,000; the project is ahead of schedule", "D", false},
			},
		},
		{
			prompt:          "Which of the following is the PRIMARY purpose of a project management plan according to PMBOK methodology?",
			domain:          "Predictive Methodologies",
			explanation:     "The project management plan integrates and consolidates all subsidiary plans and baselines. It defines how the project is executed, monitored, controlled, and closed, serving as the primary source of project information.",
			popularityScore: 2.8,
			choices: []ChoiceData{
				{"To provide detailed work instructions for team members", "A", false},
				{"To integrate and consolidate all subsidiary plans and baselines", "B", true},
				{"To serve as a contract between the project manager and sponsor", "C", false},
				{"To document all project requirements and specifications", "D", false},
			},
		},
		{
			prompt:          "In Scrum, the Product Owner is unavailable during a Sprint and cannot provide clarification on backlog items. What should the Scrum Master do?",
			domain:          "Agile Frameworks",
			explanation:     "The Scrum Master should help the team reach out to the Product Owner to ensure availability or identify a delegate, facilitating collaboration and ensuring the team has the necessary information.",
			popularityScore: 2.3,
			choices: []ChoiceData{
				{"Cancel the Sprint until the Product Owner returns", "A", false},
				{"Facilitate a discussion to ensure the Product Owner or delegate is available", "B", true},
				{"Allow the Development Team to make assumptions", "C", false},
				{"Elevate the issue to senior management immediately", "D", false},
			},
		},
		{
			prompt:          "A team practicing Kanban is experiencing long cycle times. Upon analysis, it's discovered that work items are spending too much time in the 'In Review' column. What should be the FIRST action to address this issue?",
			domain:          "Agile Frameworks",
			explanation:     "In Kanban, when a bottleneck is identified, the first step is typically to examine and potentially reduce the Work in Progress (WIP) limits for that column to force the resolution of the constraint.",
			popularityScore: 2.3,
			choices: []ChoiceData{
				{"Increase the WIP limit for the 'In Review' column", "A", false},
				{"Reduce the WIP limit for the 'In Review' column", "B", true},
				{"Skip the review process for less critical items", "C", false},
				{"Move items back to previous columns to balance the flow", "D", false},
			},
		},
		{
			prompt:          "During a Sprint Retrospective, the team identifies that they consistently struggle to complete their Sprint commitments. They often discover new requirements during development that weren't apparent during Sprint Planning. What should the team focus on improving?",
			domain:          "Agile Frameworks",
			explanation:     "This indicates issues with story analysis and decomposition during Sprint Planning. The team should focus on better requirement analysis, story decomposition, and involving the right people in planning to uncover hidden requirements.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"Reducing the number of story points committed to each Sprint", "A", false},
				{"Improving story analysis and decomposition during Sprint Planning", "B", true},
				{"Extending Sprint duration to accommodate discoveries", "C", false},
				{"Adding buffer time to each story estimate", "D", false},
			},
		},
		{
			prompt:          "During a Sprint Review, stakeholders request a significant new feature based on the product increment. How should the Scrum Team respond?",
			domain:          "Agile Frameworks",
			explanation:     "Scrum welcomes stakeholder feedback, but new work should be added to the Product Backlog, refined, and ordered by the Product Owner for future Sprints rather than being inserted mid-Sprint.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"Add the feature to the Sprint Backlog so it can start immediately", "A", false},
				{"Ask the Scrum Master to extend the Sprint to accommodate the request", "B", false},
				{"Record the request in the Product Backlog and let the Product Owner order it", "C", true},
				{"Hold an emergency Sprint Planning session with stakeholders", "D", false},
			},
		},
		{
			prompt:          "A Scrum Team’s Daily Scrum has become a status meeting directed to the Scrum Master. What is the BEST corrective action?",
			domain:          "Agile Frameworks",
			explanation:     "The Daily Scrum is a team-owned event for the Developers to inspect progress toward the Sprint Goal. The Scrum Master should coach the Developers to facilitate it themselves and refocus on planning the next 24 hours.",
			popularityScore: 2.2,
			choices: []ChoiceData{
				{"Replace the Daily Scrum with written updates to save time", "A", false},
				{"Have the Product Owner lead the Daily Scrum to keep it focused", "B", false},
				{"Coach the Developers to own the event and refocus on progress toward the Sprint Goal", "C", true},
				{"Cancel the Daily Scrum until the team plans better", "D", false},
			},
		},
		{
			prompt:          "A Kanban team notices an increasing queue in the 'Ready for Deployment' column even though deployment occurs twice a week. What should they investigate FIRST?",
			domain:          "Agile Frameworks",
			explanation:     "An accumulating queue suggests a downstream bottleneck. The team should examine whether the deployment cadence or capacity matches upstream flow before adjusting other processes.",
			popularityScore: 2.3,
			choices: []ChoiceData{
				{"Increase the WIP limits in earlier workflow stages", "A", false},
				{"Suspend upstream development until the queue clears permanently", "B", false},
				{"Review deployment capacity and cadence to address the bottleneck", "C", true},
				{"Add more detail to the column policies before changing anything", "D", false},
			},
		},
		{
			prompt:          "Which estimating technique is considered an absolute estimate in agile environments?",
			domain:          "Agile Frameworks",
			explanation:     "Absolute estimation uses fixed units like ideal hours or ideal days. Techniques such as time-based estimating give an absolute measure of effort rather than comparing items.",
			popularityScore: 2.1,
			choices: []ChoiceData{
				{"Story points using the Fibonacci sequence", "A", false},
				{"T-shirt sizing (S, M, L)", "B", false},
				{"Ideal hours estimated collaboratively", "C", true},
				{"Relative mass appraisal", "D", false},
			},
		},
		{
			prompt:          "A team wants to improve its relative estimating approach. Which practice reinforces relative estimation principles?",
			domain:          "Agile Frameworks",
			explanation:     "Relative estimation compares backlog items to one another. Techniques like Planning Poker encourage discussion and comparison, strengthening consistent sizing across the backlog.",
			popularityScore: 2.2,
			choices: []ChoiceData{
				{"Switching to individual hourly estimates", "A", false},
				{"Adopting Planning Poker with reference stories", "B", true},
				{"Estimating only after coding begins", "C", false},
				{"Removing story points and tracking only defects", "D", false},
			},
		},

		// Business Analysis Frameworks (27%)
		{
			prompt:          "A business analyst is working with stakeholders to define requirements for a new system. Several stakeholders have conflicting views on what the system should accomplish. What is the MOST effective approach to resolve these conflicts?",
			domain:          "Business Analysis",
			explanation:     "When stakeholders have conflicting requirements, the most effective approach is to facilitate collaborative workshops where stakeholders can discuss, understand each other's perspectives, and work toward consensus or compromise.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Document all requirements and let the project sponsor decide", "A", false},
				{"Facilitate collaborative workshops to reach consensus", "B", true},
				{"Implement the requirements from the most senior stakeholder", "C", false},
				{"Conduct individual interviews and create a compromise solution", "D", false},
			},
		},
		{
			prompt:          "An organization is implementing a new process that will significantly change how employees perform their daily tasks. Resistance to change is anticipated. Which change management approach would be MOST effective?",
			domain:          "Business Analysis",
			explanation:     "Effective change management involves engaging employees throughout the process, communicating benefits clearly, and providing adequate training and support. This participative approach helps reduce resistance and increases buy-in.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"Implement the change quickly to minimize disruption", "A", false},
				{"Engage employees in the change process and provide comprehensive training", "B", true},
				{"Mandate the change and provide consequences for non-compliance", "C", false},
				{"Pilot the change with a small group before full implementation", "D", false},
			},
		},
		{
			prompt:          "A project team is analyzing the business case for a proposed solution. The solution has high implementation costs but promises significant long-term benefits. The stakeholders are concerned about the initial investment. What should the business analyst emphasize when presenting the business case?",
			domain:          "Business Analysis",
			explanation:     "When dealing with high upfront costs, it's important to present the total cost of ownership and return on investment over time, showing the break-even point and long-term value creation to justify the initial investment.",
			popularityScore: 2.2,
			choices: []ChoiceData{
				{"The technical superiority of the proposed solution", "A", false},
				{"The total cost of ownership and long-term return on investment", "B", true},
				{"The risks of not implementing the solution", "C", false},
				{"The competitive advantages gained from early implementation", "D", false},
			},
		},

		// Additional challenging questions for comprehensive coverage
		{
			prompt:          "A project manager is leading a hybrid project where some components use predictive approaches while others use adaptive approaches. During execution, there's a conflict between the structured governance requirements and the need for iterative development. How should this be addressed?",
			domain:          "Project Management Fundamentals",
			explanation:     "In hybrid projects, governance should be tailored to accommodate both approaches. This might involve having structured oversight at key milestones while allowing flexibility within iterations, creating a governance framework that supports both methodologies.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Apply predictive governance to all project components", "A", false},
				{"Tailor governance to accommodate both approaches appropriately", "B", true},
				{"Use only adaptive governance throughout the project", "C", false},
				{"Separate the project into independent predictive and adaptive tracks", "D", false},
			},
		},
		{
			prompt:          "According to the PMBOK 7th Edition performance domains, which domain is PRIMARILY concerned with creating an appropriate project environment and culture?",
			domain:          "Project Management Fundamentals",
			explanation:     "The Team performance domain focuses on establishing and maintaining project culture, creating an appropriate environment, facilitating team behavior, and supporting team dynamics and leadership.",
			popularityScore: 2.3,
			choices: []ChoiceData{
				{"Stakeholder", "A", false},
				{"Team", "B", true},
				{"Development Approach and Life Cycle", "C", false},
				{"Planning", "D", false},
			},
		},
		{
			prompt:          "A project is experiencing scope creep despite having a detailed scope statement and WBS. Stakeholders keep requesting 'minor' changes that they claim won't impact the schedule or budget. What is the BEST way to address this situation?",
			domain:          "Predictive Methodologies",
			explanation:     "All changes, regardless of perceived size, should go through the formal change control process. This ensures proper evaluation of impacts and maintains project integrity. 'Minor' changes often have cumulative significant impacts.",
			popularityScore: 2.7,
			choices: []ChoiceData{
				{"Accommodate the changes since they are minor", "A", false},
				{"Implement formal change control process for all changes", "B", true},
				{"Negotiate with stakeholders to limit future changes", "C", false},
				{"Document the changes but don't process them formally", "D", false},
			},
		},
		{
			prompt:          "In an Agile project, the team has been consistently delivering working software, but the Product Owner expresses concern that the delivered features don't seem to align with the overall product vision. What should be done to address this issue?",
			domain:          "Agile Frameworks",
			explanation:     "This suggests a disconnect between the product vision and the features being developed. The team should revisit and refine the product vision collaboratively, ensure it's well-communicated, and adjust the product backlog to better align with the vision.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Continue delivering features without change", "A", false},
				{"Revisit the product vision and align backlog priorities", "B", true},
				{"Extend iterations to include more detailed documentation", "C", false},
				{"Switch to a predictive approach", "D", false},
			},
		},
		{
			prompt:          "A project team is using a Kanban system and notices that their cumulative flow diagram shows increasing work in progress across all columns. What does this indicate and what should be done?",
			domain:          "Agile Frameworks",
			explanation:     "Increasing WIP across all columns indicates that work is entering the system faster than it's being completed, leading to bottlenecks and increased cycle time. The team should limit WIP and focus on completing existing work before starting new items.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"The process is stable; continue current practices", "A", false},
				{"There are bottlenecks; limit WIP and focus on completion", "B", true},
				{"Add more team members to increase throughput", "C", false},
				{"Move items back to earlier stages for rework", "D", false},
			},
		},
		{
			prompt:          "During stakeholder analysis, a business analyst identifies a stakeholder who has high influence but low interest in the project. According to stakeholder management best practices, how should this stakeholder be managed?",
			domain:          "Business Analysis",
			explanation:     "High influence, low interest stakeholders should be kept satisfied. They have the power to impact the project but aren't actively engaged, so regular communication and ensuring their concerns are addressed is key to maintaining their support.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Monitor them with minimal effort", "A", false},
				{"Keep them satisfied through regular communication", "B", true},
				{"Engage them actively in all project decisions", "C", false},
				{"Ignore them since they have low interest", "D", false},
			},
		},
		{
			prompt:          "A project manager is working on a project where the deliverables must meet strict regulatory compliance requirements. The project team wants to use an adaptive approach, but the regulatory body requires extensive documentation and formal approval processes. What should the project manager do?",
			domain:          "Project Management Fundamentals",
			explanation:     "When regulatory requirements exist, they must be met regardless of the chosen approach. The project manager should adapt the approach to satisfy regulatory needs while retaining as much flexibility as possible within those constraints.",
			popularityScore: 2.8,
			choices: []ChoiceData{
				{"Use a purely predictive approach to meet regulatory requirements", "A", false},
				{"Adapt the approach to satisfy regulatory needs while maintaining flexibility", "B", true},
				{"Request an exemption from regulatory requirements", "C", false},
				{"Proceed with adaptive approach and handle compliance separately", "D", false},
			},
		},

		{
			prompt:          "A company has multiple related initiatives: developing a new mobile app, upgrading the website, and creating a customer loyalty program. These initiatives share resources and have interdependencies. How should these be BEST organized?",
			domain:          "Project Management Fundamentals",
			explanation:     "Related initiatives with shared resources and interdependencies should be organized as a program. Programs manage multiple related projects to achieve benefits not available when managing them individually.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"As separate independent projects", "A", false},
				{"As a program with multiple related projects", "B", true},
				{"As one large project with multiple phases", "C", false},
				{"As part of a portfolio without program management", "D", false},
			},
		},
		{
			prompt:          "What is the PRIMARY difference between a project and operations?",
			domain:          "Project Management Fundamentals",
			explanation:     "Projects are temporary with defined start and end dates to create unique deliverables, while operations are ongoing repetitive activities that sustain the business.",
			popularityScore: 2.9,
			choices: []ChoiceData{
				{"Projects are larger in scope than operations", "A", false},
				{"Projects are temporary and unique, operations are ongoing and repetitive", "B", true},
				{"Projects require more resources than operations", "C", false},
				{"Projects are more complex than operations", "D", false},
			},
		},
		{
			prompt:          "A manufacturing company runs a 24/7 production line but launches a six-month initiative to retrofit robots with new sensors. Which statement BEST classifies the work?",
			domain:          "Project Management Fundamentals",
			explanation:     "Retrofitting the robots is a temporary endeavor with a unique outcome, making it a project. The 24/7 production line remains ongoing operations because it delivers repeating value without a defined end point; the finite retrofit is what signals project work.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Both the retrofit and production line are operations", "A", false},
				{"Retrofit is a project; production line is operations", "B", true},
				{"Retrofit is operations; production line is a project", "C", false},
				{"Both are projects because robots are involved", "D", false},
			},
		},
		{
			prompt:          "During an ERP rollout, the project manager must hand over billing processes to the operations manager once the new system is live. What is the BEST indicator that ownership should transition to operations?",
			domain:          "Project Management Fundamentals",
			explanation:     "When the deliverable meets acceptance criteria and shifts into repeatable stewardship, operations should assume ownership. Handover is justified only when support processes are ready and the outcome can be sustained as business-as-usual.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"When project funding is exhausted", "A", false},
				{"When stakeholders sign acceptance and support processes are ready", "B", true},
				{"When the team feels comfortable with the system", "C", false},
				{"When scope changes stop occurring", "D", false},
			},
		},
		{
			prompt:          "Which example highlights an operations outcome rather than a project result?",
			domain:          "Project Management Fundamentals",
			explanation:     "Producing monthly financial statements is repetitive and ongoing, fitting operations because it sustains value rather than introducing change. Projects, in contrast, exist to create or modify capabilities before handing them back to operations.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Designing a new employee onboarding app", "A", false},
				{"Launching a marketing campaign for a new service", "B", false},
				{"Producing monthly financial statements", "C", true},
				{"Building a prototype for an innovative product", "D", false},
			},
		},
		{
			prompt:          "A service desk has been handling the same incident type for years but now undertakes a two-month effort to automate triage. How should the work be segmented?",
			domain:          "Project Management Fundamentals",
			explanation:     "The automation initiative is a project producing a unique capability; the ongoing incident handling remains operations. Once the one-time automation rollout is complete, stewardship of the new tooling reverts to the operational team that continues triage work.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Both activities are a single project", "A", false},
				{"Triage automation is a project, incident handling is operations", "B", true},
				{"Incident handling becomes a project because automation is involved", "C", false},
				{"Neither activity needs project management", "D", false},
			},
		},
		{
			prompt:          "Which factor MOST clearly signals the end of a project and the start of operational sustainment?",
			domain:          "Project Management Fundamentals",
			explanation:     "Formal acceptance of deliverables coupled with completed knowledge transfer indicates the project can close. Closure is driven by evidence the product meets requirements and that the receiving organization can support it, not by whether minor tasks remain on a checklist.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"The project team disbands", "A", false},
				{"Stakeholders approve deliverables and operations is ready to support them", "B", true},
				{"Budget variance reaches zero", "C", false},
				{"The sponsor schedules a lessons learned meeting", "D", false},
			},
		},
		{
			prompt:          "A retailer's e-commerce division wants to treat seasonal promotions as projects. Which criterion would justify classifying a promotion setup as a project rather than routine operations?",
			domain:          "Project Management Fundamentals",
			explanation:     "Projects deliver unique outcomes, so a one-time promotion requiring bespoke integrations is classified as project work. Routine seasonal promotions that reuse existing assets stay in operations because they do not create a unique deliverable.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"It uses the same assets as last year", "A", false},
				{"It repeats every quarter without major change", "B", false},
				{"It demands new partnerships and one-off system changes", "C", true},
				{"It is funded from the marketing budget", "D", false},
			},
		},
		{
			prompt:          "Which pair correctly matches project and operations responsibilities after a data center migration?",
			domain:          "Project Management Fundamentals",
			explanation:     "Projects execute the migration while operations maintains uptime afterwards. The project team owns time-bound change activities like cutover, whereas operations resumes responsibility for monitoring and incident response once the change is live.",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Project team monitors daily backups; operations builds risk register", "A", false},
				{"Project team executes cutover; operations manages ongoing monitoring", "B", true},
				{"Operations signs procurement contracts; project team handles ticket queues", "C", false},
				{"Operations manages stakeholder acceptance; project team owns SLAs", "D", false},
			},
		},
		{
			prompt:          "During benefits realization tracking, which metric would most likely remain under project control instead of operations?",
			domain:          "Project Management Fundamentals",
			explanation:     "Projects own benefit realization plans until transition, whereas sustaining KPIs become operations' responsibility. Metrics such as adoption rate during rollout still sit with the project because they gauge whether the change is taking hold.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Steady-state incident response time", "A", false},
				{"Adoption rate of the new solution during the transition", "B", true},
				{"Monthly recurring revenue from legacy services", "C", false},
				{"Utility costs for the existing facility", "D", false},
			},
		},
		{
			prompt:          "A PMO wants to prevent scope creep caused by operations teams submitting enhancement requests after go-live. What approach BEST maintains the project/operations boundary?",
			domain:          "Project Management Fundamentals",
			explanation:     "Establishing a change control funnel with clear criteria ensures post-go-live enhancements become separate projects when they alter scope, budget, or schedule. That governance mechanism prevents operational wish lists from silently expanding the original project.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Allow all enhancements into the existing backlog", "A", false},
				{"Create a change control intake that evaluates operational requests for new projects", "B", true},
				{"Have operations manage changes without PM oversight", "C", false},
				{"Freeze all changes after go-live", "D", false},
			},
		},
		{
			prompt:          "Which example demonstrates operations leveraging project outputs to deliver continuous value?",
			domain:          "Project Management Fundamentals",
			explanation:     "Operations uses project-created assets—such as a new CRM—to sustain service levels. After delivery, frontline teams apply the capability day to day, which is how the organization realizes the value the project created.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Project team runs the CRM indefinitely", "A", false},
				{"Customer support uses the newly implemented CRM to handle cases", "B", true},
				{"Operations designs the CRM database schema", "C", false},
				{"Stakeholders update the business case after go-live", "D", false},
			},
		},
		{
			prompt:          "A project manager discovers that a team member has been inflating their time reports. The team member explains they need the extra income due to personal financial difficulties. According to the PMI Code of Ethics, what should the project manager do?",
			domain:          "Project Management Fundamentals",
			explanation:     "The PMI Code of Ethics requires honesty and responsibility. While showing compassion for personal situations, the project manager must address the dishonest behavior and report it according to organizational policies.",
			popularityScore: 2.8,
			choices: []ChoiceData{
				{"Ignore the issue due to the team member's personal circumstances", "A", false},
				{"Address the dishonest behavior and follow organizational reporting procedures", "B", true},
				{"Allow the team member to continue but monitor them closely", "C", false},
				{"Transfer the team member to another project", "D", false},
			},
		},
		{
			prompt:          "During a vendor selection process, a supplier offers the project manager an expensive gift. According to the PMI Code of Ethics, what should the project manager do?",
			domain:          "Project Management Fundamentals",
			explanation:     "The PMI Code of Ethics prohibits accepting inappropriate gifts that could influence decision-making. The project manager should decline the gift and follow organizational policies regarding vendor relationships.",
			popularityScore: 2.7,
			choices: []ChoiceData{
				{"Accept the gift but ensure it doesn't influence the decision", "A", false},
				{"Decline the gift and follow organizational policies", "B", true},
				{"Accept the gift and declare it to the project sponsor", "C", false},
				{"Share the gift with the entire selection committee", "D", false},
			},
		},
		{
			prompt:          "During project planning, the team identifies that 'skilled developers will be available when needed.' This should be classified as a:",
			domain:          "Project Management Fundamentals",
			explanation:     "An assumption is something believed to be true but not proven. The availability of skilled developers is assumed but not guaranteed, making it an assumption that should be validated.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Risk", "A", false},
				{"Assumption", "B", true},
				{"Constraint", "C", false},
				{"Issue", "D", false},
			},
		},
		{
			prompt:          "A project team discovers that the database server is currently down and preventing development work. This should be classified as a:",
			domain:          "Project Management Fundamentals",
			explanation:     "An issue is a current problem that is impacting or will impact the project. The database server being down is happening now and preventing work, making it an issue requiring immediate attention.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"Risk", "A", false},
				{"Constraint", "B", false},
				{"Issue", "C", true},
				{"Assumption", "D", false},
			},
		},
		{
			prompt:          "During project execution, a major scope change is requested by stakeholders. Who has the PRIMARY authority to approve or reject this change?",
			domain:          "Project Management Fundamentals",
			explanation:     "The project sponsor typically has the authority to approve major changes that affect project scope, budget, or timeline. The project manager manages the change process but the sponsor makes the final decision.",
			popularityScore: 2.8,
			choices: []ChoiceData{
				{"Project manager", "A", false},
				{"Project sponsor", "B", true},
				{"Stakeholders who requested the change", "C", false},
				{"Project management office (PMO)", "D", false},
			},
		},
		{
			prompt:          "What is the PRIMARY role of a project sponsor versus a project manager?",
			domain:          "Project Management Fundamentals",
			explanation:     "The sponsor provides strategic direction and business justification while the project manager handles day-to-day execution and management. The sponsor focuses on 'what and why' while the PM focuses on 'how and when.'",
			popularityScore: 2.9,
			choices: []ChoiceData{
				{"Sponsor manages daily activities, PM provides funding", "A", false},
				{"Sponsor provides strategic direction, PM handles execution", "B", true},
				{"Both have identical responsibilities", "C", false},
				{"Sponsor handles technical decisions, PM manages stakeholders", "D", false},
			},
		},
		{
			prompt:          "A project team is demoralized after a failed sprint. As a leader, what should the project manager focus on FIRST?",
			domain:          "Project Management Fundamentals",
			explanation:     "Leadership involves inspiring and motivating people through difficult times. The PM should first acknowledge the team's feelings, help them learn from the failure, and re-energize them toward future success.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"Immediately implement process improvements to prevent future failures", "A", false},
				{"Acknowledge feelings, facilitate learning, and re-energize the team", "B", true},
				{"Report the failure to senior management with corrective actions", "C", false},
				{"Replace underperforming team members", "D", false},
			},
		},
		{
			prompt:          "High emotional intelligence (EQ) in project management is MOST valuable for:",
			domain:          "Project Management Fundamentals",
			explanation:     "Emotional intelligence is most valuable for understanding and managing relationships, reading team dynamics, and adapting communication styles to different stakeholders. It's about people skills rather than technical abilities.",
			popularityScore: 2.3,
			choices: []ChoiceData{
				{"Making technical decisions faster", "A", false},
				{"Understanding and managing stakeholder relationships", "B", true},
				{"Completing project documentation more efficiently", "C", false},
				{"Calculating project metrics more accurately", "D", false},
			},
		},

		// More Project Management Fundamentals questions
		{
			prompt:          "A project team is experiencing difficulties in decision-making due to cultural differences and varying communication styles. According to PMBOK 7th Edition principles, what should the project manager focus on to improve team effectiveness?",
			domain:          "Project Management Fundamentals",
			explanation:     "PMBOK 7th Edition emphasizes creating psychological safety and inclusive environments. The project manager should focus on building trust, encouraging diverse perspectives, and establishing clear communication protocols that respect cultural differences.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"Establish standardized communication protocols for all team members", "A", false},
				{"Create psychological safety and inclusive decision-making processes", "B", true},
				{"Assign decision-making authority to the most senior team member", "C", false},
				{"Separate team members by cultural background to reduce conflicts", "D", false},
			},
		},
		{
			prompt:          "Which of the following BEST describes the concept of 'systems thinking' as applied in PMBOK 7th Edition?",
			domain:          "Project Management Fundamentals",
			explanation:     "Systems thinking in PMBOK 7th Edition refers to understanding how project components interact with each other and with external systems, recognizing that changes in one area can impact other areas of the project or organization.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Breaking down complex problems into smaller, manageable parts", "A", false},
				{"Understanding interactions between project components and external systems", "B", true},
				{"Using systematic approaches to project planning and execution", "C", false},
				{"Implementing standardized processes across all projects", "D", false},
			},
		},

		// More Predictive Methodologies questions
		{
			prompt:          "A project manager has calculated the schedule performance index (SPI) as 0.85 and the cost performance index (CPI) as 1.15. The project sponsor asks for a forecast of the final project cost. Which formula should be used for the most accurate estimate?",
			domain:          "Predictive Methodologies",
			explanation:     "When both cost and schedule performance are expected to influence the remaining work, the most appropriate formula is EAC = AC + (BAC - EV) / (CPI × SPI). It accounts for current efficiencies in both cost and schedule as the project forecasts its final cost.",
			popularityScore: 2.1,
			choices: []ChoiceData{
				{"EAC = AC + BAC - EV", "A", false},
				{"EAC = BAC / CPI", "B", false},
				{"EAC = AC + (BAC - EV) / (CPI × SPI)", "C", true},
				{"EAC = BAC / SPI", "D", false},
			},
		},
		{
			prompt:          "During the executing process group, a project manager discovers that a key assumption documented in the project charter is no longer valid. What should be the FIRST action?",
			domain:          "Predictive Methodologies",
			explanation:     "When a key assumption becomes invalid, the first step is to assess the impact on project objectives, scope, schedule, and budget. This analysis will inform the appropriate response and potential change requests.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"Update the project charter immediately", "A", false},
				{"Assess the impact on project objectives and constraints", "B", true},
				{"Inform the project sponsor and request project termination", "C", false},
				{"Continue with the project as planned", "D", false},
			},
		},

		// More Agile Framework questions
		{
			prompt:          "A Scrum team has been working together for six months, but their velocity remains inconsistent from sprint to sprint. During retrospectives, they identify different impediments each time but don't seem to be improving overall. What should the Scrum Master focus on?",
			domain:          "Agile Frameworks",
			explanation:     "Inconsistent velocity often indicates underlying systematic issues rather than just individual impediments. The Scrum Master should focus on identifying and addressing root causes and patterns rather than just treating symptoms.",
			popularityScore: 2.3,
			choices: []ChoiceData{
				{"Extending sprint duration to allow for impediment resolution", "A", false},
				{"Identifying root causes and systemic impediments", "B", true},
				{"Reducing the amount of work committed to each sprint", "C", false},
				{"Replacing team members who are causing impediments", "D", false},
			},
		},
		{
			prompt:          "In Lean methodology, what is the primary purpose of implementing 'pull' systems in project work?",
			domain:          "Agile Frameworks",
			explanation:     "Pull systems in Lean limit work in progress and ensure that new work is only started when capacity becomes available. This optimizes flow, reduces waste, and improves quality by preventing overproduction and multitasking.",
			popularityScore: 2.0,
			choices: []ChoiceData{
				{"To increase the speed of work delivery", "A", false},
				{"To limit work in progress and optimize flow", "B", true},
				{"To ensure all team members are always busy", "C", false},
				{"To maximize resource utilization across projects", "D", false},
			},
		},

		// More Business Analysis questions
		{
			prompt:          "A business analyst is facilitating a requirements workshop where stakeholders disagree on the priority of features. The project has a fixed budget and timeline. What technique would be MOST effective for reaching consensus?",
			domain:          "Business Analysis",
			explanation:     "MoSCoW prioritization (Must have, Should have, Could have, Won't have) is particularly effective when dealing with fixed constraints as it forces stakeholders to make trade-offs and reach consensus on what's truly essential.",
			popularityScore: 2.2,
			choices: []ChoiceData{
				{"Weighted scoring model", "A", false},
				{"MoSCoW prioritization technique", "B", true},
				{"Cost-benefit analysis", "C", false},
				{"Stakeholder voting", "D", false},
			},
		},
		{
			prompt:          "An organization is experiencing resistance to a new system implementation. Users complain that the system doesn't match their current workflow. What should the business analyst recommend as the BEST approach?",
			domain:          "Business Analysis",
			explanation:     "When there's a mismatch between system capabilities and current workflows, the best approach is to analyze both the current and future state processes, then determine the optimal combination of system configuration and process changes.",
			popularityScore: 2.4,
			choices: []ChoiceData{
				{"Modify the system to match current workflows exactly", "A", false},
				{"Analyze current vs. future state processes and optimize both", "B", true},
				{"Force users to adapt to the new system without changes", "C", false},
				{"Implement the system in phases to gradually introduce changes", "D", false},
			},
		},

		// Advanced scenario-based questions
		{
			prompt:          "A project manager is managing a complex project with multiple external vendors. Two vendors have dependencies between their deliverables, but they are reluctant to coordinate directly due to competitive concerns. How should the project manager handle this situation?",
			domain:          "Project Management Fundamentals",
			explanation:     "In complex multi-vendor situations, the project manager should act as an intermediary to facilitate coordination while protecting each vendor's competitive interests. This involves managing interfaces and dependencies without compromising confidentiality.",
			popularityScore: 2.7,
			choices: []ChoiceData{
				{"Require vendors to work together despite competitive concerns", "A", false},
				{"Act as intermediary to facilitate coordination while protecting interests", "B", true},
				{"Eliminate dependencies by redesigning the project scope", "C", false},
				{"Allow each vendor to work independently and resolve conflicts later", "D", false},
			},
		},
		{
			prompt:          "During project execution, a critical team member informs the project manager that they will be unavailable for the next two weeks due to a family emergency. This absence will impact the critical path. What should the project manager do FIRST?",
			domain:          "Predictive Methodologies",
			explanation:     "The first step is to assess the specific impact on the critical path and project schedule. This analysis will inform the best response strategy, whether it's resource reallocation, schedule adjustment, or other mitigation measures.",
			popularityScore: 2.5,
			choices: []ChoiceData{
				{"Immediately find a replacement team member", "A", false},
				{"Assess the impact on the critical path and schedule", "B", true},
				{"Inform the sponsor and request a schedule extension", "C", false},
				{"Redistribute the work among other team members", "D", false},
			},
		},
		{
			prompt:          "A product owner in an Agile project keeps changing priorities between sprints, causing confusion and frustration among the development team. The Scrum Master should:",
			domain:          "Agile Frameworks",
			explanation:     "The Scrum Master should coach the Product Owner on effective backlog management and help establish criteria for priority changes. This addresses the root cause while maintaining the Product Owner's authority over backlog prioritization.",
			popularityScore: 2.6,
			choices: []ChoiceData{
				{"Ask the team to ignore priority changes during sprints", "A", false},
				{"Coach the Product Owner on backlog management and priority stability", "B", true},
				{"Escalate the issue to management for resolution", "C", false},
				{"Implement a change control process for backlog modifications", "D", false},
			},
		},
		{
			prompt:          "A business analyst discovers that a requested feature will require significant changes to the underlying data architecture. The development team estimates this will triple the implementation time. What should the business analyst do?",
			domain:          "Business Analysis",
			explanation:     "When significant impacts are discovered, the business analyst should present alternatives with their trade-offs to stakeholders. This might include alternative solutions, phased implementation, or accepting the extended timeline based on business value.",
			popularityScore: 2.3,
			choices: []ChoiceData{
				{"Recommend removing the feature from the project scope", "A", false},
				{"Present alternatives and trade-offs to stakeholders for decision", "B", true},
				{"Ask the development team to find a faster implementation approach", "C", false},
				{"Proceed with the feature as requested despite the impact", "D", false},
			},
		},

		{
			prompt:          "Which statement BEST describes the purpose of a project charter?",
			domain:          "Project Management Fundamentals",
			explanation:     "The charter formally authorizes the project, links it to organizational strategy, and grants the project manager authority. It documents the business justification and clearly states who is empowered to lead the work.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"It details the WBS and activity durations", "A", false},
				{"It authorizes the project and assigns the project manager", "B", true},
				{"It lists all operational procedures", "C", false},
				{"It defines the risk register in full", "D", false},
			},
		},
		{
			prompt:          "In an adaptive life cycle, which set of stages typically repeats within each iteration?",
			domain:          "Agile Frameworks",
			explanation:     "Adaptive cycles emphasize an iteration of planning, executing, reviewing, and adapting. Each sprint or increment revisits that loop so the team can respond to feedback and evolving requirements.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Initiate, authorize, close", "A", false},
				{"Plan, execute, review, adjust", "B", true},
				{"Procure, negotiate, mobilize", "C", false},
				{"Design, install, decommission", "D", false},
			},
		},
		{
			prompt:          "The CAPM Exam Content Outline groups questions into domains. What is the PRIMARY value of these domains for exam preparation?",
			domain:          "Project Management Fundamentals",
			explanation:     "The Exam Content Outline clarifies domain weightings so candidates allocate study time intentionally, ensuring preparation aligns with how PMI scores the exam.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"They replace the PMBOK Guide completely", "A", false},
				{"They show weighting so you can prioritize study topics", "B", true},
				{"They are optional reading suggestions", "C", false},
				{"They list every formula you must memorize", "D", false},
			},
		},
		{
			prompt:          "A project team keeps accepting 'minor' feature tweaks without evaluating impact, and the schedule is slipping. Which action BEST controls scope creep?",
			domain:          "Predictive Methodologies",
			explanation:     "Running all changes through integrated change control prevents uncontrolled scope growth by forcing impact analysis and formal approval before work begins.",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Document tweaks but skip approval", "A", false},
				{"Route every change through the formal change control process", "B", true},
				{"Ask the team to work faster", "C", false},
				{"Ignore stakeholder requests", "D", false},
			},
		},
		{
			prompt:          "Which output of risk management provides the foundation for assigning owners and budgets to risk responses?",
			domain:          "Predictive Methodologies",
			explanation:     "The risk management plan defines categories, roles, timing, and funding approaches for risk work, giving the team a framework before they begin identifying specific risks.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Risk register", "A", false},
				{"Risk management plan", "B", true},
				{"Lessons learned register", "C", false},
				{"Issue log", "D", false},
			},
		},
		{
			prompt:          "During Identify Risks, which technique MOST effectively captures a broad set of threats and opportunities early in the project?",
			domain:          "Predictive Methodologies",
			explanation:     "Facilitated brainstorming with subject-matter experts surfaces diverse risk perspectives quickly because it encourages cross-functional input before analysis narrows the list.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Quantitative simulation", "A", false},
				{"Brainstorming with SMEs", "B", true},
				{"Reserve analysis", "C", false},
				{"Variance analysis", "D", false},
			},
		},
		{
			prompt:          "Midway through execution, risk audits reveal several high-priority responses losing effectiveness. What should the project manager do FIRST?",
			domain:          "Predictive Methodologies",
			explanation:     "Reassessing the risk strategy and adjusting response plans within Monitor Risks keeps the strategy alive, ensuring responses remain effective as conditions change.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Close the risks immediately", "A", false},
				{"Update the risk response strategies and communicate changes", "B", true},
				{"Transfer all risks to the sponsor", "C", false},
				{"Ignore the audit because responses are in place", "D", false},
			},
		},
		{
			prompt:          "Which option correctly differentiates a project, program, and portfolio?",
			domain:          "Project Management Fundamentals",
			explanation:     "Projects deliver unique outputs, programs manage related projects to realize broader benefits, and portfolios align projects and programs with strategic objectives. Each level operates at a different horizon of value delivery.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Project manages multiple portfolios", "A", false},
				{"Program groups related projects; portfolio aligns strategic investments", "B", true},
				{"Portfolio delivers a single unique product", "C", false},
				{"Project and program are identical", "D", false},
			},
		},
		{
			prompt:          "A schedule is behind and the PM overlaps phases previously planned in sequence to regain time. Which technique is being applied?",
			domain:          "Predictive Methodologies",
			explanation:     "Fast tracking runs activities in parallel to compress the schedule without adding resources, accepting increased coordination risk in exchange for time savings.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Resource leveling", "A", false},
				{"Fast tracking", "B", true},
				{"Crashing", "C", false},
				{"Decomposition", "D", false},
			},
		},
		{
			prompt:          "To meet a milestone, the PM adds experienced contractors to a critical path activity, increasing cost but shortening duration. Which schedule compression method is this?",
			domain:          "Predictive Methodologies",
			explanation:     "Crashing adds resources to decrease duration at higher cost, trading budget for schedule improvement on critical path activities.",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Crashing", "A", true},
				{"Fast tracking", "B", false},
				{"Monte Carlo", "C", false},
				{"Rolling wave", "D", false},
			},
		},

		{
			prompt:          "Which type of PMO primarily provides templates, best practices, and training while leaving project control with project managers?",
			domain:          "Project Management Fundamentals",
			explanation:     "A supportive PMO offers guidance, templates, and training without exerting direct control, acting more like an internal consultancy than a governing body.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Directive", "A", false},
				{"Supportive", "B", true},
				{"Controlling", "C", false},
				{"Hybrid", "D", false},
			},
		},

		{
			prompt:          "What is the PRIMARY role of a Project Management Office (PMO)?",
			domain:          "Project Management Fundamentals",
			explanation:     "PMOs standardize governance, provide oversight, and align projects with organizational strategy, ensuring consistent delivery practices across initiatives.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Provide day-to-day team supervision", "A", false},
				{"Deliver operational outputs", "B", false},
				{"Standardize project practices and ensure strategic alignment", "C", true},
				{"Decide stakeholder priorities independently", "D", false},
			},
		},

		{
			prompt:          "Which of the following is an example of an Enterprise Environmental Factor (EEF)?",
			domain:          "Project Management Fundamentals",
			explanation:     "Enterprise Environmental Factors are conditions outside the project team's control—such as organizational culture or regulatory requirements—that influence how the project is executed.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Lessons learned database", "A", false},
				{"Quality policy set by industry regulators", "B", true},
				{"Risk register", "C", false},
				{"Change control template", "D", false},
			},
		},

		{
			prompt:          "Which item is classified as an Organizational Process Asset (OPA)?",
			domain:          "Project Management Fundamentals",
			explanation:     "Organizational Process Assets are internal resources such as templates, procedures, and historical records that teams reuse to accelerate project work.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Government regulations", "A", false},
				{"Company project templates", "B", true},
				{"Market conditions", "C", false},
				{"Currency exchange rates", "D", false},
			},
		},

		{
			prompt:          "In a functional organizational structure, who typically has the most authority over resources?",
			domain:          "Project Management Fundamentals",
			explanation:     "In a functional structure, resource authority sits with functional managers, leaving project managers to negotiate for people and assets on a case-by-case basis.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Project manager", "A", false},
				{"Functional manager", "B", true},
				{"Sponsor", "C", false},
				{"Scrum Master", "D", false},
			},
		},

		{
			prompt:          "Which statement BEST distinguishes a functional manager from a project manager?",
			domain:          "Project Management Fundamentals",
			explanation:     "Functional managers focus on ongoing departmental operations, whereas project managers lead temporary initiatives that introduce change before handing results back to the line organization.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Functional managers own temporary goals", "A", false},
				{"Project managers handle operational staffing", "B", false},
				{"Functional managers optimize departmental operations; project managers deliver temporary change", "C", true},
				{"Project managers report to functional staff", "D", false},
			},
		},

		{
			prompt:          "In Scrum, which role ensures the team adheres to Scrum values and removes impediments?",
			domain:          "Agile Frameworks",
			explanation:     "The Scrum Master facilitates events, coaches the team on agile principles, and shields members from distractions so they can focus on delivering increments.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Product Owner", "A", false},
				{"Scrum Master", "B", true},
				{"Project Sponsor", "C", false},
				{"Stakeholder", "D", false},
			},
		},

		{
			prompt:          "Which is NOT a Scrum artifact?",
			domain:          "Agile Frameworks",
			explanation:     "Scrum artifacts are the Product Backlog, Sprint Backlog, and Increment—items that capture intended work and value produced—so a burn-down chart doesn't qualify.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Product Backlog", "A", false},
				{"Sprint Burn-down Chart", "B", true},
				{"Increment", "C", false},
				{"Sprint Backlog", "D", false},
			},
		},

		{
			prompt:          "Which event in Scrum allows the team to inspect its process and plan improvements?",
			domain:          "Agile Frameworks",
			explanation:     "The Sprint Retrospective is the dedicated forum for the team to inspect its process, discuss what helped or hurt delivery, and plan targeted improvements.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Sprint Planning", "A", false},
				{"Daily Scrum", "B", false},
				{"Sprint Review", "C", false},
				{"Sprint Retrospective", "D", true},
			},
		},

		{
			prompt:          "Which pair correctly identifies a positive risk (opportunity) and a negative risk (threat)?",
			domain:          "Predictive Methodologies",
			explanation:     "Delivering early by fast-tracking is an opportunity because it can be exploited for added benefit, whereas supplier bankruptcy is a threat that requires mitigation or contingency planning.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Supplier bankruptcy / deliver earlier", "A", false},
				{"Deliver earlier / supplier bankruptcy", "B", true},
				{"Quality issue / budget surplus", "C", false},
				{"Cost overrun / resource added", "D", false},
			},
		},

		{
			prompt:          "Which process balances cost and quality trade-offs to meet stakeholder expectations?",
			domain:          "Project Management Fundamentals",
			explanation:     "Cost and quality trade-offs are addressed in Plan Quality Management, where the team defines standards, metrics, and responsibilities before execution begins.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Plan Quality Management", "A", true},
				{"Control Costs", "B", false},
				{"Manage Stakeholder Engagement", "C", false},
				{"Monitor Communications", "D", false},
			},
		},
		{
			prompt:          "Which approach suits projects with well-understood requirements and stable scope?",
			domain:          "Predictive Methodologies",
			explanation:     "Predictive approaches work best when scope is fixed and requirements are clear because detailed upfront planning can proceed without constant rework, unlike adaptive approaches that embrace change.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Adaptive", "A", false},
				{"Predictive", "B", true},
				{"Hybrid", "C", false},
				{"Incremental", "D", false},
			},
		},
		{
			prompt:          "What does Minimum Viable Product (MVP) represent in agile delivery?",
			domain:          "Agile Frameworks",
			explanation:     "A Minimum Viable Product is the smallest release that delivers genuine customer value and invites feedback, allowing the team to validate assumptions before investing further.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"A prototype with no user testing", "A", false},
				{"The smallest value-bearing release for feedback", "B", true},
				{"A complete product", "C", false},
				{"Internal documentation", "D", false},
			},
		},
		{
			prompt:          "In product development, which concept focuses on delivering the smallest increment that generates measurable benefit?",
			domain:          "Agile Frameworks",
			explanation:     "A Minimum Business Increment delivers a measurable business outcome—something leadership can evaluate for tangible impact—whereas an MVP primarily validates viability.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Minimal Business Integration", "A", false},
				{"Minimum Business Increment", "B", true},
				{"Migration Batch Item", "C", false},
				{"Measured Business Input", "D", false},
			},
		},
		{
			prompt:          "Which process defines quality policies, metrics, and responsibilities for the project?",
			domain:          "Predictive Methodologies",
			explanation:     "Plan Quality Management sets the quality standards, metrics, and responsibilities that will guide later execution and control activities.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Plan Quality Management", "A", true},
				{"Control Quality", "B", false},
				{"Manage Quality", "C", false},
				{"Close Project", "D", false},
			},
		},
		{
			prompt:          "Scope planning primarily results in which key deliverable?",
			domain:          "Project Management Fundamentals",
			explanation:     "Scope planning develops both the scope management plan and the scope baseline, providing the blueprint for how scope will be defined, validated, and controlled.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Risk management plan", "A", false},
				{"Scope management plan and scope baseline", "B", true},
				{"Resource breakdown structure", "C", false},
				{"Procurement statement of work", "D", false},
			},
		},
		{
			prompt:          "Which components make up a performance measurement baseline?",
			domain:          "Predictive Methodologies",
			explanation:     "The scope, schedule, and cost baselines collectively form the performance measurement baseline, locking in the targets used to track project performance.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Risk, issue, change baselines", "A", false},
				{"Scope, schedule, cost baselines", "B", true},
				{"Quality, resource, procurement baselines", "C", false},
				{"Stakeholder and communication baselines", "D", false},
			},
		},
		{
			prompt:          "During Collect Requirements, which technique ensures stakeholder needs are documented clearly?",
			domain:          "Project Management Fundamentals",
			explanation:     "Stakeholder interviews permit detailed exploration of expectations and uncover nuanced needs that might be missed by broad surveys or workshops.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Fast tracking", "A", false},
				{"Stakeholder interviews", "B", true},
				{"Lead time analysis", "C", false},
				{"Decision tree analysis", "D", false},
			},
		},
		{
			prompt:          "Which development approach creates an early model to gather feedback before building the final product?",
			domain:          "Project Management Fundamentals",
			explanation:     "Prototyping helps stakeholders visualize solutions early, gather concrete feedback, and refine requirements before committing to full-scale build.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Monte Carlo simulation", "A", false},
				{"Benchmarking", "B", false},
				{"Prototyping", "C", true},
				{"Decomposition", "D", false},
			},
		},
		{
			prompt:          "Which statement BEST differentiates a work breakdown structure (WBS) from WBS activities?",
			domain:          "Predictive Methodologies",
			explanation:     "The WBS decomposes deliverables into smaller components, while the subsequent Define Activities process turns those work packages into schedule tasks; they are related but distinct steps.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"WBS lists time-phased tasks", "A", false},
				{"WBS decomposes scope; activities emerge during Define Activities", "B", true},
				{"Activities precede WBS creation", "C", false},
				{"WBS is only used in agile", "D", false},
			},
		},

		{
			prompt:          "Which communication method pushes information to stakeholders when it is not time-critical?",
			domain:          "Project Management Fundamentals",
			explanation:     "Push communication—such as emails or memos—sends information to recipients without confirming receipt, essentially broadcasting updates for later review.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Push communication", "A", true},
				{"Pull communication", "B", false},
				{"Interactive communication", "C", false},
				{"Passive listening", "D", false},
			},
		},
		{
			prompt:          "Which communication skill BEST supports building trust and uncovering hidden concerns?",
			domain:          "Project Management Fundamentals",
			explanation:     "Active listening demonstrates empathy, encourages stakeholders to share concerns, and often reveals expectations that might otherwise stay hidden.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Speaking louder", "A", false},
				{"Active listening", "B", true},
				{"Sending status reports", "C", false},
				{"Using technical jargon", "D", false},
			},
		},
		{
			prompt:          "Which factor can distort a message between sender and receiver in project communications?",
			domain:          "Project Management Fundamentals",
			explanation:     "Noise—such as cultural bias, jargon, or language barriers—can distort messages between sender and receiver, so project managers must surface and address those filters.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"A well-written email", "A", false},
				{"Active listening", "B", false},
				{"Cultural/linguistic noise", "C", true},
				{"Face-to-face discussion", "D", false},
			},
		},

		{
			prompt:          "Your project planned value (PV) at month 4 is $400,000. Earned value (EV) is $360,000 and actual cost (AC) is $420,000. Based on CPI and SPI, what best describes performance?",
			domain:          "Predictive Methodologies",
			explanation:     "CPI = EV/AC = 0.86 (over budget) and SPI = EV/PV = 0.90 (behind schedule). Values below 1.0 directly signal cost overrun and schedule slippage, respectively.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Ahead of schedule, over budget", "A", false},
				{"Behind schedule, under budget", "B", false},
				{"Over budget and behind schedule", "C", true},
				{"Under budget and on schedule", "D", false},
			},
		},
		{
			prompt:          "A project has EV $550,000, AC $500,000 and PV $600,000. What is the cost variance (CV) and interpretation?",
			domain:          "Predictive Methodologies",
			explanation:     "CV = EV - AC = $50,000, a positive variance that indicates the project is currently under budget because earned value exceeds actual spending.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"CV = -$50,000, over budget", "A", false},
				{"CV = $50,000, under budget", "B", true},
				{"CV = $100,000, ahead of schedule", "C", false},
				{"CV = -$100,000, behind schedule", "D", false},
			},
		},
		{
			prompt:          "At completion of month 6, EV is $720,000, AC is $800,000, BAC is $1,200,000. Management expects current cost efficiency to continue. Which EAC formula applies?",
			domain:          "Predictive Methodologies",
			explanation:     "If cost performance (CPI) is expected to continue, the appropriate forecast is EAC = BAC / CPI. With CPI at 0.90, the projected total cost rises to about $1.33M.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"EAC = AC + (BAC - EV)", "A", false},
				{"EAC = AC + (BAC - EV)/(CPI * SPI)", "B", false},
				{"EAC = BAC / CPI", "C", true},
				{"EAC = BAC - CV", "D", false},
			},
		},
		{
			prompt:          "Given BAC $900,000 and EAC recalculated at $1,050,000, what is variance at completion (VAC) and meaning?",
			domain:          "Predictive Methodologies",
			explanation:     "VAC = BAC - EAC = -$150,000, so the project is tracking toward a cost overrun of that amount if performance trends persist.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"VAC = $150,000, underrun expected", "A", false},
				{"VAC = -$150,000, overrun expected", "B", true},
				{"VAC = $0, on budget", "C", false},
				{"VAC cannot be calculated without AC", "D", false},
			},
		},
		{
			prompt:          "At 40% complete, CPI is 0.92 and AC is $460,000. What is EV?",
			domain:          "Predictive Methodologies",
			explanation:     "Because CPI = EV/AC, you can rearrange to EV = CPI × AC = 0.92 × 460,000 = $423,200, showing how earned value is derived from known metrics.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"$423,200", "A", true},
				{"$460,000", "B", false},
				{"$500,000", "C", false},
				{"$420,000", "D", false},
			},
		},
		{
			prompt:          "With BAC $2,000,000, EV $1,100,000, AC $1,300,000 and remaining work budget $900,000, management wants to know the to-complete performance index (TCPI) to meet BAC. What is TCPI and implication?",
			domain:          "Predictive Methodologies",
			explanation:     "TCPI(BAC) = (BAC - EV)/(BAC - AC) = (2,000,000 - 1,100,000)/(2,000,000 - 1,300,000) = 900,000/700,000 = 1.29, indicating future cost efficiency must exceed current performance because the value is above 1.",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"1.29, efficiency must improve", "A", true},
				{"0.71, efficiency can decrease", "B", false},
				{"1.00, maintain current performance", "C", false},
				{"Cannot be determined", "D", false},
			},
		},
		{
			prompt:          "EV is $375,000 and PV is $420,000. What is schedule variance and interpretation?",
			domain:          "Predictive Methodologies",
			explanation:     "SV = EV - PV = -$45,000, so the team has delivered $45,000 less value than planned at this point, confirming a schedule lag in earned value terms.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"SV = $45,000, ahead of schedule", "A", false},
				{"SV = -$45,000, behind schedule", "B", true},
				{"SV = $0, on schedule", "C", false},
				{"SV = $-375,000, over budget", "D", false},
			},
		},
		{
			prompt:          "If revised estimates suggest remaining work requires $350,000 using bottom-up analysis, which ETC formula applies regardless of CPI?",
			domain:          "Predictive Methodologies",
			explanation:     "A bottom-up Estimate to Complete replaces the remaining portion of the baseline with a fresh estimate created from detailed analysis of the unfinished work.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"ETC = (BAC - EV)/CPI", "A", false},
				{"ETC = BAC - EV", "B", false},
				{"ETC = $350,000", "C", true},
				{"ETC = AC + EV", "D", false},
			},
		},
		{
			prompt:          "EV $800,000, AC $820,000, CPI trending upward for three months. How should the PM interpret the trend?",
			domain:          "Predictive Methodologies",
			explanation:     "An improving CPI trend suggests cost performance is stabilizing, so the project manager should continue monitoring rather than overreacting to a single data point.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Ignore trend; it is meaningless", "A", false},
				{"Recognize improving cost efficiency and update forecasts", "B", true},
				{"Assume cost overrun will worsen", "C", false},
				{"Cancel the project", "D", false},
			},
		},
		{
			prompt:          "SPI is 0.88 while CPI is 1.05. What combination of schedule and cost status does this reveal?",
			domain:          "Predictive Methodologies",
			explanation:     "SPI below 1 signals schedule delay while CPI above 1 shows the project remains under budget; the two indices must be interpreted separately to understand the full performance picture.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Behind schedule, under budget", "A", true},
				{"Ahead of schedule, over budget", "B", false},
				{"Behind schedule, over budget", "C", false},
				{"Ahead of schedule, under budget", "D", false},
			},
		},

		{
			prompt:          "What is the primary purpose of the Plan Risk Management process?",
			domain:          "Predictive Methodologies",
			explanation:     "Plan Risk Management defines how risk activities will be conducted and tailored to the project, effectively planning the risk approach before diving into identification.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"To list every risk", "A", false},
				{"To specify how risk management will be performed", "B", true},
				{"To calculate contingency reserves", "C", false},
				{"To update the stakeholder register", "D", false},
			},
		},
		{
			prompt:          "Which document captures roles, budgeting, and definitions for risk categories?",
			domain:          "Predictive Methodologies",
			explanation:     "The risk management plan captures methodology, roles, funding, timing, and categories, serving as the playbook for every subsequent risk process.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Risk register", "A", false},
				{"Risk management plan", "B", true},
				{"Lessons learned register", "C", false},
				{"Issue log", "D", false},
			},
		},
		{
			prompt:          "Which of the following BEST describes the contents of a project management plan?",
			domain:          "Predictive Methodologies",
			explanation:     "The project management plan integrates all subsidiary plans—scope, schedule, cost, quality, resource, communications, risk, procurement, stakeholder—and establishes baselines and tailoring decisions so the team knows how work will be executed and controlled.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"It lists only the project schedule and budget", "A", false},
				{"It consolidates baselines and all subsidiary management plans", "B", true},
				{"It contains solely stakeholder contact information", "C", false},
				{"It is limited to the risk register and issue log", "D", false},
			},
		},
		{
			prompt:          "During Identify Risks, what key output records individual risks and their characteristics?",
			domain:          "Predictive Methodologies",
			explanation:     "The risk register catalogs each identified risk along with triggers, owners, probability, impact, and potential responses for easy reference throughout the project.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Risk report", "A", false},
				{"Risk register", "B", true},
				{"Risk breakdown structure", "C", false},
				{"Project charter", "D", false},
			},
		},
		{
			prompt:          "Qualitative Risk Analysis primarily helps the team do what?",
			domain:          "Predictive Methodologies",
			explanation:     "Perform Qualitative Risk Analysis prioritizes risks for further action by assessing the probability and impact scores, helping the team focus on the most significant threats and opportunities.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Compute monetary exposure", "A", false},
				{"Rank risks by probability and impact", "B", true},
				{"Create contingency reserves", "C", false},
				{"Close low priority risks", "D", false},
			},
		},
		{
			prompt:          "Which technique assesses the reliability of risk data before prioritizing?",
			domain:          "Predictive Methodologies",
			explanation:     "Risk data quality assessments evaluate whether information is complete, accurate, and consistent so that later analysis rests on trustworthy inputs.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Sensitivity analysis", "A", false},
				{"Risk data quality assessment", "B", true},
				{"SWOT analysis", "C", false},
				{"Expected monetary value", "D", false},
			},
		},
		{
			prompt:          "Which quantitative risk analysis tool simulates project outcomes using random variables?",
			domain:          "Predictive Methodologies",
			explanation:     "Monte Carlo simulation models probability distributions through many iterations to forecast potential cost or schedule outcomes and quantify overall uncertainty.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Probability impact matrix", "A", false},
				{"Monte Carlo simulation", "B", true},
				{"Risk urgency assessment", "C", false},
				{"Delphi technique", "D", false},
			},
		},
		{
			prompt:          "A well-written risk statement typically follows which structure?",
			domain:          "Predictive Methodologies",
			explanation:     "Well-formed risk statements follow an IF [cause] THEN [effect] structure so the team sees both the trigger and the potential consequence in a single sentence.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"IF cause THEN effect", "A", true},
				{"Effect THEN opportunity", "B", false},
				{"Risk equals mitigation", "C", false},
				{"Issue THEN workaround", "D", false},
			},
		},
		{
			prompt:          "Risk categories are often organized using what tool to ensure comprehensive coverage?",
			domain:          "Predictive Methodologies",
			explanation:     "A risk breakdown structure organizes risks by source, creating a hierarchical “family tree” that highlights concentration areas like technical, external, or organizational risks.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Risk register", "A", false},
				{"Risk breakdown structure", "B", true},
				{"Work breakdown structure", "C", false},
				{"Responsibility assignment matrix", "D", false},
			},
		},
		{
			prompt:          "Assigning an owner to each major risk occurs in which process?",
			domain:          "Predictive Methodologies",
			explanation:     "Plan Risk Responses designates specific owners to implement agreed responses, ensuring accountability rather than leaving actions unassigned.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Identify Risks", "A", false},
				{"Perform Qualitative Risk Analysis", "B", false},
				{"Plan Risk Responses", "C", true},
				{"Monitor Risks", "D", false},
			},
		},
		{
			prompt:          "Which risk response strategy for threats involves ceasing the risky activity entirely?",
			domain:          "Predictive Methodologies",
			explanation:     "Avoidance eliminates the threat by changing the plan so the risky situation no longer exists—essentially removing exposure altogether.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Mitigate", "A", false},
				{"Transfer", "B", false},
				{"Avoid", "C", true},
				{"Accept", "D", false},
			},
		},
		{
			prompt:          "Contingency reserves are typically created to handle which type of risks?",
			domain:          "Predictive Methodologies",
			explanation:     "Contingency reserves cover identified risks with planned responses, whereas management reserves address unknown-unknowns that fall outside the risk register.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Unknown unknowns", "A", false},
				{"Known risks with planned responses", "B", true},
				{"Scope creep", "C", false},
				{"Defect repairs", "D", false},
			},
		},
		{
			prompt:          "Implement Risk Responses ensures what?",
			domain:          "Predictive Methodologies",
			explanation:     "Implement Risk Responses ensures the planned actions are carried out and assessed for effectiveness, translating risk strategy into execution.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Risks are prioritized", "A", false},
				{"Responses are carried out", "B", true},
				{"Reserves are released", "C", false},
				{"Lessons are archived", "D", false},
			},
		},
		{
			prompt:          "Which process keeps risk responses current and monitors residual and secondary risks?",
			domain:          "Predictive Methodologies",
			explanation:     "Monitor Risks continually tracks identified, residual, and secondary risks while checking whether response strategies and audits remain effective.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Implement Risk Responses", "A", false},
				{"Monitor Risks", "B", true},
				{"Identify Risks", "C", false},
				{"Perform Quantitative Risk Analysis", "D", false},
			},
		},
		{
			prompt:          "A probability and impact matrix is primarily used during which process?",
			domain:          "Predictive Methodologies",
			explanation:     "The probability-and-impact matrix supports Perform Qualitative Risk Analysis by ranking risks so the team knows which warrant deeper analysis or immediate action.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Plan Risk Management", "A", false},
				{"Perform Qualitative Risk Analysis", "B", true},
				{"Perform Quantitative Risk Analysis", "C", false},
				{"Plan Risk Responses", "D", false},
			},
		},
		{
			prompt:          "When implementing a mitigation response, a new risk emerges because of the action. What type of risk is this?",
			domain:          "Predictive Methodologies",
			explanation:     "Secondary risks arise as a direct result of implementing responses; teams must identify and plan for them whenever they adjust the original strategy.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"Residual risk", "A", false},
				{"Secondary risk", "B", true},
				{"Contingent risk", "C", false},
				{"Unknown risk", "D", false},
			},
		},
		{
			prompt:          "Risk audits are conducted in Monitor Risks to achieve what objective?",
			domain:          "Predictive Methodologies",
			explanation:     "Risk audits evaluate how well the risk management process and individual responses are working, offering input for improvements.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Create risk breakdown structures", "A", false},
				{"Assess response effectiveness", "B", true},
				{"Generate new risk categories", "C", false},
				{"Close all risks", "D", false},
			},
		},
		{
			prompt:          "Why is early risk identification beneficial?",
			domain:          "Predictive Methodologies",
			explanation:     "Identifying risks early gives the team time to craft effective responses and reduces the likelihood of being blindsided later in the project.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"It increases contingency reserves", "A", false},
				{"It allows more time to plan responses", "B", true},
				{"It eliminates all threats", "C", false},
				{"It delays stakeholder input", "D", false},
			},
		},
		{
			prompt:          "Which strategy is appropriate for opportunities and aims to ensure the opportunity happens?",
			domain:          "Predictive Methodologies",
			explanation:     "Exploiting an opportunity involves changing the plan to guarantee the upside occurs, rather than simply hoping to capture it.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Mitigate", "A", false},
				{"Accept", "B", false},
				{"Exploit", "C", true},
				{"Transfer", "D", false},
			},
		},
		{
			prompt:          "What distinguishes a residual risk from other risk types?",
			domain:          "Predictive Methodologies",
			explanation:     "Residual risks are the leftover exposure that persists even after response plans have been executed, and they must be tracked accordingly.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"It arises from responses", "A", false},
				{"It exists after responses and is accepted", "B", true},
				{"It is unknown", "C", false},
				{"It cannot be documented", "D", false},
			},
		},
		{
			prompt:          "Risk thresholds define what?",
			domain:          "Predictive Methodologies",
			explanation:     "Risk thresholds specify how much risk the organization is willing to tolerate before additional action is required, providing clear triggers for escalation.",
			popularityScore: 3.1,
			choices: []ChoiceData{
				{"The list of all risks", "A", false},
				{"The acceptable level of risk exposure", "B", true},
				{"A qualitative ranking", "C", false},
				{"The contingency reserve amount", "D", false},
			},
		},
		{
			prompt:          "Which document summarizes overall project risk exposure and high-level trends for stakeholders?",
			domain:          "Predictive Methodologies",
			explanation:     "The risk report consolidates overall risk status, trends, and summary information, giving stakeholders a big-picture view of exposure and response effectiveness.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Risk register", "A", false},
				{"Risk report", "B", true},
				{"Risk breakdown structure", "C", false},
				{"Lessons learned register", "D", false},
			},
		},
		{
			prompt:          "When a risk is transferred to a third party, what should the project team still do?",
			domain:          "Predictive Methodologies",
			explanation:     "After transferring a risk contractually, the project manager still monitors the arrangement to confirm the response is effective, because accountability for oversight remains with the project.",
			popularityScore: 3.0,
			choices: []ChoiceData{
				{"Ignore the risk", "A", false},
				{"Monitor the transfer agreement and outcomes", "B", true},
				{"Close the risk immediately", "C", false},
				{"Reassign it to the sponsor", "D", false},
			},
		},
		// Hard Question Drills
		{
			prompt:          "A project is halfway through development when the founding sponsor resigns and a new executive inherits accountability.\nThe incoming sponsor demands a comprehensive re-evaluation of the business case before authorising more spend.\nTeam members fear a stop-work order will demoralise them and derail regulatory milestones.\nFunctional leads want direction while the PMO emphasises disciplined governance.\nWhat should the project manager do first?",
			domain:          "Hard Question",
			explanation:     "Meeting the new sponsor to reconfirm objectives and routing any adjustments through change control sustains alignment without stalling delivery.",
			popularityScore: 3.6,
			choices: []ChoiceData{
				{"Pause all delivery until a revised charter is issued to guarantee legitimacy", "A", false},
				{"Prepare a defensive justification deck and ask the PMO to endorse the original plan", "B", false},
				{"Engage the new sponsor to validate objectives and document required updates through formal change control", "C", true},
				{"Continue execution unchanged while gathering informal feedback to avoid delays", "D", false},
			},
		},
		{
			prompt:          "A capital program shows EV $4.2M, AC $5.1M, and three consecutive months of declining CPI just as the steering committee requests a decision briefing.\nDirectors are anxious about sunk costs but want a storyline that focuses on next steps rather than a formula tutorial.\nSeveral senior stakeholders wonder whether the program should pivot or be re-scoped before it consumes more funding.\nYou need to present the performance data without eroding confidence.\nHow should the update be positioned?",
			domain:          "Hard Question",
			explanation:     "Framing the CPI erosion, presenting response options, and recommending scope reprioritisation with refreshed benefits keeps leaders focused on value recovery.",
			popularityScore: 3.6,
			choices: []ChoiceData{
				{"Insist the baseline remain untouched because investments are already committed", "A", false},
				{"Recommend terminating immediately without analysing strategic alternatives", "B", false},
				{"Promise to rebaseline cost and schedule while leaving scope and benefits unchanged", "C", false},
				{"Explain the CPI trend, outline viable choices, and propose reprioritising scope with a refreshed benefit case", "D", true},
			},
		},
		{
			prompt:          "During UAT a client insists on adding a predictive portfolio module before launch because competitors tout similar features.\nThe longtime sponsor resigned yesterday, leaving an interim sponsor who wants no surprises at next week's board meeting.\nYour senior integration developer begins a four-week approved holiday after this sprint and is the only person who understands the new data pipelines.\nCompliance deadlines are fixed and the delivery team fears derailment.\nHow should the project manager respond?",
			domain:          "Hard Question",
			explanation:     "Re-engaging governance, sequencing the request through change control, and capturing knowledge before the holiday protects alignment and schedule.",
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"Accept the module immediately and mandate overtime so the release date stays intact", "A", false},
				{"Delay the release until the developer returns while avoiding the interim sponsor", "B", false},
				{"Realign objectives with the interim sponsor, process the request formally, stage delivery after knowledge handoff, and brief the board on the updated plan", "C", true},
				{"Decline the new request outright and defer it to the next planning cycle", "D", false},
			},
		},
		{
			prompt:          "Three sprints remain before a regulated launch when the client's CFO demands embedded ESG dashboards for an investor roadshow.\nThe original sponsor has just left the company and the acting sponsor expects a confident board message within ten days.\nLegal warns that changing scope without analysis would breach the master services agreement.\nYour only data architect has a prepaid six-week sabbatical that begins in two weeks.\nWhich response best protects delivery credibility?",
			domain:          "Hard Question",
			explanation:     "Running expedited change control with interim leadership, phasing dashboards post-launch, and capturing knowledge before the sabbatical balances responsiveness with contractual guardrails.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Guarantee the dashboards on the original date and ask the architect to cancel the sabbatical", "A", false},
				{"Let the team rearrange work informally and notify governance later", "B", false},
				{"Facilitate a governance session with the acting sponsor and legal, process the dashboards through change control, stage delivery after knowledge handoff, and brief the board on the revised forecast", "C", true},
				{"Pause the entire release until a permanent sponsor is appointed and the architect returns", "D", false},
			},
		},
		{
			prompt:          "A regional rollout is mid-integration when the client requests a loyalty analytics layer to debut at an upcoming marketing summit.\nThe regional sponsor has been promoted and wants the enhancement showcased at their first steering committee.\nYour analytics lead, who owns the cross-market data flows, starts approved family leave in three weeks.\nRegulatory filings are due shortly after go-live, increasing anxiety about delays.\nWhat should the project manager do?",
			domain:          "Hard Question",
			explanation:     "Assessing impact with the new sponsor, routing the request through change control, and completing knowledge transfer before leave keeps compliance and stakeholder confidence intact.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Absorb the loyalty scope immediately and extend team hours to keep pace", "A", false},
				{"Freeze the rollout until the analytics lead returns", "B", false},
				{"Meet the new sponsor to evaluate the change, run formal impact analysis, plan phased delivery after the handoff, and protect regulatory milestones", "C", true},
				{"Tell the sponsor the enhancement is impossible within the current wave", "D", false},
			},
		},
	}
	questions = append(questions, earnedValueDrillQuestions()...)
	questions = append(questions, pertDrillQuestions()...)
	questions = append(questions, projectOperationsClassificationQuestions()...)
	questions = append(questions, stakeholderSalienceDrillQuestions()...)
	questions = append(questions, teamMotivationDrillQuestions()...)

	return questions
}

func earnedValueDrillQuestions() []QuestionData {
	const domain = "Earned Value Drill"

	questions := make([]QuestionData, 0, 50)

	rotation := 0

	pvScenarios := []struct {
		name    string
		bac     float64
		planned float64
	}{
		{"Aurora", 120000, 0.35},
		{"Beacon", 90000, 0.50},
		{"Catalyst", 150000, 0.32},
		{"Delta", 78000, 0.62},
		{"Equinox", 210000, 0.40},
	}

	for i, s := range pvScenarios {
		pv := round2(s.bac * s.planned)
		distractors := currencyDistractors(pv)
		explanation := fmt.Sprintf("PV = BAC × planned %% complete = %s × %s = %s.", formatCurrency(s.bac), formatPercent(s.planned), formatCurrency(pv))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[PV] Project %s has a BAC of %s and is planned to be %s complete at this checkpoint. What is the planned value (PV)?", s.name, formatCurrency(s.bac), formatPercent(s.planned)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.4,
			choices:         createChoices(pv, distractors, rotation+i, formatCurrency),
		})
	}

	evScenarios := []struct {
		name   string
		bac    float64
		actual float64
	}{
		{"Fusion", 130000, 0.42},
		{"Glacier", 175000, 0.36},
		{"Halcyon", 95000, 0.58},
		{"Ion", 160000, 0.47},
		{"Juniper", 205000, 0.33},
	}

	for i, s := range evScenarios {
		ev := round2(s.bac * s.actual)
		distractors := currencyDistractors(ev)
		explanation := fmt.Sprintf("EV = BAC × actual %% complete = %s × %s = %s.", formatCurrency(s.bac), formatPercent(s.actual), formatCurrency(ev))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[EV] Project %s has a BAC of %s and is actually %s complete. What is the earned value (EV)?", s.name, formatCurrency(s.bac), formatPercent(s.actual)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.4,
			choices:         createChoices(ev, distractors, rotation+len(pvScenarios)+i, formatCurrency),
		})
	}

	acScenarios := []struct {
		name       string
		components []struct {
			label string
			value float64
		}
	}{
		{
			name: "Ignite",
			components: []struct {
				label string
				value float64
			}{
				{"Requirements", 9500},
				{"Design", 14000},
				{"Testing", 7200},
			},
		},
		{
			name: "Kestrel",
			components: []struct {
				label string
				value float64
			}{
				{"Analysis", 11800},
				{"Development", 25750},
				{"Quality Assurance", 9800},
			},
		},
		{
			name: "Lumen",
			components: []struct {
				label string
				value float64
			}{
				{"Hardware", 16500},
				{"Software", 19400},
				{"Training", 8600},
			},
		},
		{
			name: "Meridian",
			components: []struct {
				label string
				value float64
			}{
				{"Sprint 1", 11250},
				{"Sprint 2", 12400},
				{"Infrastructure", 10200},
			},
		},
		{
			name: "Nova",
			components: []struct {
				label string
				value float64
			}{
				{"Installation", 8700},
				{"Configuration", 15450},
				{"User Support", 9300},
			},
		},
	}

	acRotationBase := rotation + len(pvScenarios) + len(evScenarios)
	for i, s := range acScenarios {
		var total float64
		parts := make([]string, 0, len(s.components))
		for _, cmp := range s.components {
			total += cmp.value
			parts = append(parts, fmt.Sprintf("%s %s", cmp.label, formatCurrency(cmp.value)))
		}
		total = round2(total)
		distractors := currencyDistractors(total)
		prompt := fmt.Sprintf("[AC] Project %s has incurred the following costs to date: %s. What is the Actual Cost (AC)?", s.name, strings.Join(parts, "; "))
		explanation := fmt.Sprintf("AC = sum of actual costs = %s.", formatCurrency(total))
		questions = append(questions, QuestionData{
			prompt:          prompt,
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.3,
			choices:         createChoices(total, distractors, acRotationBase+i, formatCurrency),
		})
	}

	metricScenarios := []struct {
		name    string
		bac     float64
		planned float64
		actual  float64
		ac      float64
	}{
		{"Horizon", 180000, 0.45, 0.40, 85000},
		{"Lighthouse", 220000, 0.50, 0.48, 102000},
		{"Momentum", 160000, 0.38, 0.42, 69000},
		{"Nebula", 195000, 0.60, 0.55, 112000},
		{"Odyssey", 250000, 0.52, 0.47, 118000},
		{"Pioneer", 140000, 0.44, 0.40, 62000},
	}

	svRotationBase := acRotationBase + len(acScenarios)
	for i := 0; i < 5 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		pv := round2(s.bac * s.planned)
		ev := round2(s.bac * s.actual)
		sv := round2(ev - pv)
		distractors := currencyDistractors(sv)
		status := "behind schedule"
		if sv > 0 {
			status = "ahead of schedule"
		} else if sv == 0 {
			status = "exactly on schedule"
		}
		prompt := fmt.Sprintf("[SV] Project %s has EV = %s and PV = %s. What is the schedule variance (SV)?", s.name, formatCurrency(ev), formatCurrency(pv))
		explanation := fmt.Sprintf("SV = EV - PV = %s - %s = %s (%s).", formatCurrency(ev), formatCurrency(pv), formatCurrency(sv), status)
		questions = append(questions, QuestionData{
			prompt:          prompt,
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.5,
			choices:         createChoices(sv, distractors, svRotationBase+i, formatCurrency),
		})
	}

	cvRotationBase := svRotationBase + 5
	for i := 0; i < 5 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		ev := round2(s.bac * s.actual)
		cv := round2(ev - s.ac)
		distractors := currencyDistractors(cv)
		status := "over budget"
		if cv > 0 {
			status = "under budget"
		} else if cv == 0 {
			status = "on budget"
		}
		prompt := fmt.Sprintf("[CV] Project %s has EV = %s and AC = %s. What is the cost variance (CV)?", s.name, formatCurrency(ev), formatCurrency(s.ac))
		explanation := fmt.Sprintf("CV = EV - AC = %s - %s = %s (%s).", formatCurrency(ev), formatCurrency(s.ac), formatCurrency(cv), status)
		questions = append(questions, QuestionData{
			prompt:          prompt,
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.5,
			choices:         createChoices(cv, distractors, cvRotationBase+i, formatCurrency),
		})
	}

	spiRotationBase := cvRotationBase + 5
	for i := 0; i < 5 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		pv := round2(s.bac * s.planned)
		ev := round2(s.bac * s.actual)
		spi := round2(ev / pv)
		distractors := ratioDistractors(spi)
		explanation := fmt.Sprintf("SPI = EV / PV = %s / %s = %s.", formatCurrency(ev), formatCurrency(pv), formatRatio(spi))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[SPI] Project %s reports EV = %s and PV = %s. What is the schedule performance index (SPI)?", s.name, formatCurrency(ev), formatCurrency(pv)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.6,
			choices:         createChoices(spi, distractors, spiRotationBase+i, formatRatio),
		})
	}

	cpiRotationBase := spiRotationBase + 5
	for i := 0; i < 5 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		ev := round2(s.bac * s.actual)
		cpi := round2(ev / s.ac)
		distractors := ratioDistractors(cpi)
		explanation := fmt.Sprintf("CPI = EV / AC = %s / %s = %s.", formatCurrency(ev), formatCurrency(s.ac), formatRatio(cpi))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[CPI] Project %s has EV = %s and AC = %s. What is the cost performance index (CPI)?", s.name, formatCurrency(ev), formatCurrency(s.ac)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.6,
			choices:         createChoices(cpi, distractors, cpiRotationBase+i, formatRatio),
		})
	}

	eacCpiRotationBase := cpiRotationBase + 5
	for i := 0; i < 5 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		ev := round2(s.bac * s.actual)
		cpi := ev / s.ac
		eac := round2(s.bac / cpi)
		distractors := currencyDistractors(eac)
		explanation := fmt.Sprintf("Assuming CPI remains constant, EAC = BAC / CPI = %s / %s = %s.", formatCurrency(s.bac), formatRatio(round2(cpi)), formatCurrency(eac))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[EAC (CPI)] Project %s has BAC = %s, EV = %s, and AC = %s. Assuming cost performance stays the same, what is the estimate at completion (EAC)?", s.name, formatCurrency(s.bac), formatCurrency(ev), formatCurrency(s.ac)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.6,
			choices:         createChoices(eac, distractors, eacCpiRotationBase+i, formatCurrency),
		})
	}

	eacOneRotationBase := eacCpiRotationBase + 5
	for i := 0; i < 3 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		ev := round2(s.bac * s.actual)
		eac := round2(s.ac + (s.bac - ev))
		distractors := currencyDistractors(eac)
		explanation := fmt.Sprintf("Assuming remaining work follows the original plan, EAC = AC + (BAC - EV) = %s + (%s - %s) = %s.", formatCurrency(s.ac), formatCurrency(s.bac), formatCurrency(ev), formatCurrency(eac))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[EAC (One-Time)] Project %s experienced a one-time cost spike. Given BAC = %s, EV = %s, and AC = %s, what is the estimate at completion (EAC) if future work proceeds as planned?", s.name, formatCurrency(s.bac), formatCurrency(ev), formatCurrency(s.ac)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.5,
			choices:         createChoices(eac, distractors, eacOneRotationBase+i, formatCurrency),
		})
	}

	eacCombinedRotationBase := eacOneRotationBase + 3
	for i := 0; i < 2 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		pv := round2(s.bac * s.planned)
		ev := round2(s.bac * s.actual)
		cpi := ev / s.ac
		spi := ev / pv
		eac := round2(s.ac + (s.bac-ev)/(cpi*spi))
		distractors := currencyDistractors(eac)
		explanation := fmt.Sprintf("Using CPI and SPI, EAC = AC + (BAC - EV)/(CPI × SPI) = %s + (%s - %s)/(%.2f × %.2f) = %s.", formatCurrency(s.ac), formatCurrency(s.bac), formatCurrency(ev), round2(cpi), round2(spi), formatCurrency(eac))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[EAC (CPI & SPI)] Project %s has BAC = %s, EV = %s, AC = %s, CPI = %.2f, and SPI = %.2f. Using both indices, what is the estimate at completion (EAC)?", s.name, formatCurrency(s.bac), formatCurrency(ev), formatCurrency(s.ac), round2(cpi), round2(spi)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.5,
			choices:         createChoices(eac, distractors, eacCombinedRotationBase+i, formatCurrency),
		})
	}

	etcRotationBase := eacCombinedRotationBase + 2
	for i := 0; i < 3 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		ev := round2(s.bac * s.actual)
		cpi := ev / s.ac
		eac := round2(s.bac / cpi)
		etc := round2(eac - s.ac)
		distractors := currencyDistractors(etc)
		explanation := fmt.Sprintf("With CPI held constant, ETC = EAC - AC = %s - %s = %s.", formatCurrency(eac), formatCurrency(s.ac), formatCurrency(etc))
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[ETC] Project %s assumes future work will follow current cost efficiency (CPI). Given BAC = %s, EV = %s, and AC = %s, what is the estimate to complete (ETC)?", s.name, formatCurrency(s.bac), formatCurrency(ev), formatCurrency(s.ac)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.4,
			choices:         createChoices(etc, distractors, etcRotationBase+i, formatCurrency),
		})
	}

	vacRotationBase := etcRotationBase + 3
	for i := 0; i < 2 && i < len(metricScenarios); i++ {
		s := metricScenarios[i]
		ev := round2(s.bac * s.actual)
		cpi := ev / s.ac
		eac := round2(s.bac / cpi)
		vac := round2(s.bac - eac)
		distractors := currencyDistractors(vac)
		status := "overrun"
		if vac > 0 {
			status = "underrun"
		} else if vac == 0 {
			status = "on target"
		}
		prompt := fmt.Sprintf("[VAC] Project %s expects CPI to hold. With BAC = %s and EAC = %s, what is the variance at completion (VAC)?", s.name, formatCurrency(s.bac), formatCurrency(eac))
		explanation := fmt.Sprintf("VAC = BAC - EAC = %s - %s = %s (%s).", formatCurrency(s.bac), formatCurrency(eac), formatCurrency(vac), status)
		questions = append(questions, QuestionData{
			prompt:          prompt,
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.4,
			choices:         createChoices(vac, distractors, vacRotationBase+i, formatCurrency),
		})
	}

	statusScenarios := []struct {
		name string
		ev   float64
		pv   float64
		ac   float64
	}{
		{"Quartz", 95000, 100000, 90000},
		{"Solstice", 132500, 128000, 141000},
		{"AuroraPrime", 108000, 102000, 97000},
		{"BeaconRise", 95000, 90000, 98000},
		{"CascadeTrail", 88000, 94000, 82000},
		{"DynamoEdge", 76000, 81000, 84000},
		{"EmberGlow", 124500, 118000, 117000},
		{"Fathom", 112200, 110000, 120000},
		{"GlacierPoint", 67500, 72000, 64000},
		{"HarborLine", 54000, 58500, 56000},
		{"IonStorm", 152000, 147000, 140000},
		{"JadeVista", 91500, 87400, 94000},
		{"LumenTrail", 119800, 125600, 127400},
		{"MesaRidge", 101400, 98600, 97800},
		{"NimbusGate", 87300, 86000, 90200},
		{"OrchidBay", 76500, 80200, 72000},
		{"PinnaclePeak", 99000, 105000, 102500},
		{"QuantumLoop", 142300, 137000, 133800},
		{"Riverstone", 128900, 131500, 129600},
		{"SummitForge", 156400, 149500, 162800},
		{"TundraField", 87200, 91000, 82800},
		{"Ultraviolet", 118600, 112400, 125900},
	}

	for _, s := range statusScenarios {
		ev := round2(s.ev)
		pv := round2(s.pv)
		ac := round2(s.ac)
		sv := round2(ev - pv)
		cv := round2(ev - ac)

		scheduleStatus := "on schedule"
		if sv > 0 {
			scheduleStatus = "ahead of schedule"
		} else if sv < 0 {
			scheduleStatus = "behind schedule"
		}

		costStatus := "on budget"
		if cv > 0 {
			costStatus = "under budget"
		} else if cv < 0 {
			costStatus = "over budget"
		}

		explanation := fmt.Sprintf("SV = EV - PV = %s - %s = %s (%s). CV = EV - AC = %s - %s = %s (%s).", formatCurrency(ev), formatCurrency(pv), formatCurrency(sv), scheduleStatus, formatCurrency(ev), formatCurrency(ac), formatCurrency(cv), costStatus)

		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[Status] Project %s reports EV = %s, PV = %s, and AC = %s. Which option best describes its performance?", s.name, formatCurrency(ev), formatCurrency(pv), formatCurrency(ac)),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"Ahead of schedule and under budget", "A", scheduleStatus == "ahead of schedule" && costStatus == "under budget"},
				{"Ahead of schedule and over budget", "B", scheduleStatus == "ahead of schedule" && costStatus == "over budget"},
				{"Behind schedule and under budget", "C", scheduleStatus == "behind schedule" && costStatus == "under budget"},
				{"Behind schedule and over budget", "D", scheduleStatus == "behind schedule" && costStatus == "over budget"},
			},
		})
	}

	return questions
}

func projectOperationsClassificationQuestions() []QuestionData {
	const domain = "Project Operations Classification Drill"

	options := []struct {
		key  string
		text string
	}{
		{"Project", "Project – temporary endeavor delivering a unique outcome"},
		{"Program", "Program – coordinated management of related projects and supporting work"},
		{"Portfolio", "Portfolio – collection of work aligned to strategic objectives"},
		{"Operations", "Operations – ongoing activities that sustain business-as-usual"},
	}

	items := []struct {
		prompt      string
		correct     string
		explanation string
	}{
		{
			prompt:      "A regional bank hires you to implement a new core banking module with a defined go-live and handover to support.",
			correct:     "Project",
			explanation: "The work is temporary, unique, and produces a specified deliverable before closing, satisfying the definition of a project.",
		},
		{
			prompt:      "The service desk rotates analysts to resolve password resets, hardware issues, and small enhancements day after day.",
			correct:     "Operations",
			explanation: "This is repeatable, ongoing work that maintains existing services, which characterizes operations.",
		},
		{
			prompt:      "A transformation director coordinates CRM replacement, data cleansing, and change management projects to deliver a unified customer view.",
			correct:     "Program",
			explanation: "Multiple related projects are governed together to realize shared benefits, making this a program.",
		},
		{
			prompt:      "An investment council reviews and balances all technology initiatives to ensure the mix best advances corporate strategy.",
			correct:     "Portfolio",
			explanation: "Selecting and optimizing the collection of projects and programs for strategy alignment is portfolio management.",
		},
		{
			prompt:      "Marketing assembles a cross-functional team to design and launch a one-time product awareness campaign for a seasonal release.",
			correct:     "Project",
			explanation: "A time-bound effort producing a unique campaign deliverable fits the definition of a project.",
		},
		{
			prompt:      "A semiconductor plant runs fabrication lines continuously, producing the same chip design under standard work instructions.",
			correct:     "Operations",
			explanation: "The activity is ongoing and repetitive to sustain output, so it is operations work.",
		},
		{
			prompt:      "An agile release train synchronizes multiple teams building related features and manages shared risks and benefits together.",
			correct:     "Program",
			explanation: "Coordinated governance of interdependent projects to deliver combined capabilities defines a program.",
		},
		{
			prompt:      "A public-sector PMO chooses which construction, IT, and policy projects receive funding based on mission alignment and risk appetite.",
			correct:     "Portfolio",
			explanation: "Balancing investment choices across diverse work for strategic fit exemplifies portfolio management.",
		},
		{
			prompt:      "An IT department performs a one-time migration of legacy applications into a new cloud tenancy, then disbands the team.",
			correct:     "Project",
			explanation: "It is a temporary effort with a unique outcome and closure, which is characteristic of a project.",
		},
		{
			prompt:      "Facilities management schedules routine cleaning, safety inspections, and equipment checks to keep offices open year-round.",
			correct:     "Operations",
			explanation: "This work repeats indefinitely to sustain services, so it is operational in nature.",
		},
		{
			prompt:      "A customer experience initiative bundles contact center modernization, knowledge base redesign, and training workstreams under one governance board.",
			correct:     "Program",
			explanation: "Related projects are coordinated to deliver a combined improvement, indicating a program.",
		},
		{
			prompt:      "A nonprofit executive team weighs grant proposals, sunsets low-value work, and ensures resources back the most critical causes.",
			correct:     "Portfolio",
			explanation: "They manage the mix of initiatives to achieve strategic objectives, which is portfolio stewardship.",
		},
		{
			prompt:      "A utility company constructs a new solar farm from permitting through commissioning before turning it over to operations.",
			correct:     "Project",
			explanation: "Delivering a unique asset with a clear end state is project work.",
		},
		{
			prompt:      "The HR department administers onboarding, benefits enrollment, and exit processing using standard templates and SLAs every week.",
			correct:     "Operations",
			explanation: "These tasks are ongoing, standardized, and repetitive, so they fall under operations.",
		},
		{
			prompt:      "A cybersecurity uplift effort coordinates vulnerability remediation, awareness campaigns, and tooling deployments to realize an enterprise risk reduction target.",
			correct:     "Program",
			explanation: "Multiple related initiatives are orchestrated to achieve collective benefits, matching the definition of a program.",
		},
		{
			prompt:      "A technology steering committee evaluates all digital investments, sequenced roadmaps, and funding requests for balanced strategic coverage.",
			correct:     "Portfolio",
			explanation: "They govern the slate of work for strategic alignment and value optimization, which is portfolio management.",
		},
		{
			prompt:      "An instructional design team builds a new certification prep course with kickoff, development sprints, pilot delivery, and closeout.",
			correct:     "Project",
			explanation: "The effort is temporary and produces a unique training asset, indicating project work.",
		},
		{
			prompt:      "A fulfillment center executes picking, packing, and shipping activities daily based on established takt times and KPIs.",
			correct:     "Operations",
			explanation: "The workflow is continuous and repeatable to sustain product delivery, which is operations.",
		},
		{
			prompt:      "A new product introduction office integrates design, manufacturing ramp-up, supplier readiness, and launch communications under one governance framework.",
			correct:     "Program",
			explanation: "Coordinating several related projects for combined launch benefits aligns with program management.",
		},
		{
			prompt:      "A healthcare network aligns telehealth, analytics, and patient engagement investments to achieve enterprise transformation goals.",
			correct:     "Portfolio",
			explanation: "They curate the collection of initiatives to realize strategic outcomes, demonstrating portfolio management.",
		},
	}

	questions := make([]QuestionData, 0, len(items))
	for i, item := range items {
		rotation := i % len(options)
		choices := make([]ChoiceData, len(options))
		for j := range options {
			idx := (rotation + j) % len(options)
			option := options[idx]
			choices[j] = ChoiceData{
				text:      option.text,
				label:     string(rune('A' + j)),
				isCorrect: option.key == item.correct,
			}
		}

		questions = append(questions, QuestionData{
			prompt:          item.prompt,
			domain:          domain,
			explanation:     item.explanation,
			popularityScore: 3.5,
			choices:         choices,
		})
	}

	return questions
}

func stakeholderSalienceDrillQuestions() []QuestionData {
	const domain = "Stakeholder Salience Drill"

	items := []struct {
		prompt      string
		explanation string
		choices     []ChoiceData
	}{
		{
			prompt:      "A cybersecurity vendor retains crucial penetration-test authority on your hybrid program. They have formal legitimacy through the contract and can delay go-live approvals, but they rarely demand immediate action. How should they be classified in the salience model?",
			explanation: "Legitimate and powerful stakeholders without urgent claims are dominant; they require routine engagement and formal governance touchpoints.",
			choices: []ChoiceData{
				{"Dormant", "A", false},
				{"Dominant", "B", true},
				{"Discretionary", "C", false},
				{"Dangerous", "D", false},
			},
		},
		{
			prompt:      "An influential regulator threatens injunctions if remediation milestones slip. The agency has statutory legitimacy, can halt operations, and demands an immediate update. Which salience category fits best?",
			explanation: "Stakeholders with legitimate authority, coercive power, and urgent claims are definitive; they must receive priority attention and decisions.",
			choices: []ChoiceData{
				{"Definitive", "A", true},
				{"Dependent", "B", false},
				{"Dangerous", "C", false},
				{"Latent", "D", false},
			},
		},
		{
			prompt:      "A community advocacy group organizes media coverage about accessibility gaps in your digital transformation. They lack contractual legitimacy but mobilize public attention quickly. Which attribute set drives their salience?",
			explanation: "Without formal legitimacy yet demonstrating urgency and the ability to disrupt via public pressure, they are dangerous stakeholders.",
			choices: []ChoiceData{
				{"Dangerous (power + urgency)", "A", true},
				{"Discretionary (legitimacy only)", "B", false},
				{"Dependent (legitimacy + urgency)", "C", false},
				{"Dominant (power + legitimacy)", "D", false},
			},
		},
		{
			prompt:      "Your CFO sponsor requests an expedited forecast workshop but states it can happen anytime this quarter. They control portfolio funding and have chartered the initiative. How would you respond based on salience?",
			explanation: "With power and legitimacy but no urgent claim, the sponsor remains dominant; schedule the engagement promptly but within planned governance cadence.",
			choices: []ChoiceData{
				{"Treat as definitive: escalate immediately", "A", false},
				{"Treat as dominant: prioritize within existing cadence", "B", true},
				{"Treat as discretionary: optional communication", "C", false},
				{"Treat as dependent: wait for urgency", "D", false},
			},
		},
		{
			prompt:      "A user group holds legitimate representation through the change advisory board and constantly raises urgent usability defects, yet they cannot enforce decisions alone. Which salience category applies?",
			explanation: "Legitimacy plus urgency without controlling power places them in the dependent category; empower them via alliances to convert needs into action.",
			choices: []ChoiceData{
				{"Dependent", "A", true},
				{"Discretionary", "B", false},
				{"Dormant", "C", false},
				{"Dangerous", "D", false},
			},
		},
		{
			prompt:      "Legacy system engineers hold deep architecture knowledge but no formal role in the new program. They are willing to advise when asked. How should you categorize them?",
			explanation: "Stakeholders with legitimacy only—recognized expertise but no authority or urgent claim—are discretionary; engage them through advisory channels.",
			choices: []ChoiceData{
				{"Discretionary", "A", true},
				{"Dormant", "B", false},
				{"Dependent", "C", false},
				{"Definitive", "D", false},
			},
		},
		{
			prompt:      "A senior architect threatens to block release approvals unless a backlog item is prioritized. They lack delegated authority but can delay quality sign-off. Urgency is high. How do you treat them?",
			explanation: "Without legitimate charter but wielding power and urgency, they are dangerous. Mitigate through escalation paths and transparent governance.",
			choices: []ChoiceData{
				{"Dangerous", "A", true},
				{"Dormant", "B", false},
				{"Dominant", "C", false},
				{"Definitive", "D", false},
			},
		},
		{
			prompt:      "An executive sponsor is reassigned. The new leader has not yet engaged and shows little urgency, but retains full decision authority. How does their salience shift?",
			explanation: "They still possess legitimacy and power, keeping them dominant despite low urgency; swiftly onboard them to maintain commitment.",
			choices: []ChoiceData{
				{"Dominant", "A", true},
				{"Dormant", "B", false},
				{"Definitive", "C", false},
				{"Dependent", "D", false},
			},
		},
		{
			prompt:      "A legal advisor is on retainer but has no pressing issues and minimal interaction with the program. They remain available if required. Which classification fits?",
			explanation: "Stakeholders with dormant power—authority but lacking legitimacy or urgency—are dormant; monitor them and activate when necessary.",
			choices: []ChoiceData{
				{"Dormant", "A", true},
				{"Discretionary", "B", false},
				{"Dangerous", "C", false},
				{"Dependent", "D", false},
			},
		},
		{
			prompt:      "A customer advisory council shares strategic insights and insists on a rapid response to new market regulations. They lack enforcement power but influence roadmap legitimacy and urgency. How should they be managed?",
			explanation: "Legitimacy combined with urgent claims places them in the dependent category; provide high-touch engagement and align with authoritative sponsors.",
			choices: []ChoiceData{
				{"Dependent", "A", true},
				{"Discretionary", "B", false},
				{"Dominant", "C", false},
				{"Dangerous", "D", false},
			},
		},
	}

	questions := make([]QuestionData, 0, len(items))
	for _, item := range items {
		questions = append(questions, QuestionData{
			prompt:          item.prompt,
			domain:          domain,
			explanation:     item.explanation,
			popularityScore: 3.6,
			choices:         item.choices,
		})
	}

	return questions
}

func teamMotivationDrillQuestions() []QuestionData {
	questions := []QuestionData{
		{
			prompt: `A newly merged project team has just been introduced after a restructuring.
Members speak politely, avoid disagreement, and repeatedly look to the project manager for instruction.
When asked to volunteer for backlog work, people hesitate because trust is not yet established.
Stakeholders are watching closely because the initiative is politically visible.
Which Tuckman stage best describes the situation?`,
			domain:          "Team Motivation Drill",
			explanation:     "Polite interactions, reliance on the leader, and uncertainty about roles indicate the Forming stage of Tuckman's model.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Forming", "A", true},
				{"Storming", "B", false},
				{"Norming", "C", false},
				{"Performing", "D", false},
			},
		},
		{
			prompt: `A cross-functional agile squad is in the second sprint together and disagreements erupt in daily syncs.
Developers challenge estimates, testers question acceptance criteria, and the product owner is frustrated by perceived negativity.
Despite the tension, members are beginning to surface real issues that were previously hidden.
Executives fear the volatility means the team is failing.
Which interpretation of Tuckman's stages should the PM offer?`,
			domain:          "Team Motivation Drill",
			explanation:     "Open conflict and testing of boundaries are hallmarks of the Storming stage; surfacing issues is a healthy step toward cohesion.",
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"The team has regressed to Forming and needs to restart chartering", "A", false},
				{"Storming is expected and signals that the team is working through power and process questions", "B", true},
				{"They are already Performing and should increase velocity targets", "C", false},
				{"They are in Adjourning and should prepare handover", "D", false},
			},
		},
		{
			prompt: `After months of collaboration, a project team anticipates go-live and shares lessons in a retrospective.
Members recognise each other's contributions, propose future experiments, and talk about how to celebrate outcomes.
Some teammates feel anxious about being assigned elsewhere once delivery is complete.
The sponsor wants the PM to manage morale during the transition.
Which Tuckman stage is the team entering?`,
			domain:          "Team Motivation Drill",
			explanation:     "Closure activities and concerns about what comes next signal the Adjourning stage of team development.",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Norming", "A", false},
				{"Performing", "B", false},
				{"Adjourning", "C", true},
				{"Storming", "D", false},
			},
		},
		{
			prompt: `Two product teams have merged mid-release.
The first week together is quiet, the second week becomes confrontational about coding standards, and now facilitators notice steady collaboration, peer coaching, and shared ownership of impediments.
Velocity is stabilising and retrospectives surface continuous improvement ideas.
What Tuckman stage does this pattern illustrate?`,
			domain:          "Team Motivation Drill",
			explanation:     "After progressing through conflict, the team is demonstrating Norming behaviours with shared norms and mutual support.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Forming", "A", false},
				{"Storming", "B", false},
				{"Norming", "C", true},
				{"Performing", "D", false},
			},
		},
		{
			prompt: `A high-performing DevOps group works autonomously, swaps roles as needed, and resolves production incidents without escalation.
Stakeholders trust their forecasts and the team self-corrects when experiments fail.
The PM notices conflict is channelled into solutioning rather than personal criticism.
Which Tuckman stage most accurately describes the team?`,
			domain:          "Team Motivation Drill",
			explanation:     "Self-managing behaviour, constructive conflict, and consistent delivery are characteristics of the Performing stage.",
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"Storming", "A", false},
				{"Norming", "B", false},
				{"Performing", "C", true},
				{"Adjourning", "D", false},
			},
		},
		{
			prompt: `During one-on-ones several engineers say they cannot focus on a complex migration because layoff rumours dominate conversation.
They are less worried about recognition than about paying rent if the rumour proves true.
The sponsor wants to motivate the team with gamified leaderboards.
Which layer of Maslow's hierarchy must the PM address first?`,
			domain:          "Team Motivation Drill",
			explanation:     "Fear about job security reflects Maslow's Safety needs, which must be stabilised before higher-level motivators resonate.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Self-actualisation", "A", false},
				{"Esteem", "B", false},
				{"Safety", "C", true},
				{"Social", "D", false},
			},
		},
		{
			prompt: `A business analyst requests training budget and a rotational assignment to work on strategic initiatives.
The analyst already has strong peer relationships, a competitive salary, and public recognition from leadership.
They explain they need to feel they are growing their full potential.
Which Maslow level best captures the unmet need?`,
			domain:          "Team Motivation Drill",
			explanation:     "Seeking growth and fulfilling potential aligns with the Self-actualisation level in Maslow's hierarchy.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Esteem", "A", false},
				{"Self-actualisation", "B", true},
				{"Social", "C", false},
				{"Safety", "D", false},
			},
		},
		{
			prompt: `A tester enjoys the team dynamic but quietly mentions they feel invisible compared to engineers who demo new features.
They already receive market pay and have time to take breaks, yet they crave acknowledgement in sprint reviews.
Which Maslow need is most pressing?`,
			domain:          "Team Motivation Drill",
			explanation:     "Desiring appreciation and status indicates the Esteem level in Maslow's hierarchy of needs.",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Physiological", "A", false},
				{"Safety", "B", false},
				{"Social", "C", false},
				{"Esteem", "D", true},
			},
		},
		{
			prompt: `A delivery centre updates its HR policies to provide reliable equipment, ergonomic seating, and clear overtime rules.
Management assumes this will dramatically increase intrinsic motivation, yet engagement scores barely change.
Which Herzberg concept best explains the outcome?`,
			domain:          "Team Motivation Drill",
			explanation:     "Policies and working conditions are hygiene factors; improving them prevents dissatisfaction but does not create lasting motivation.",
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"Achievement motivators", "A", false},
				{"Hygiene factors reduce dissatisfaction but do not energise performance", "B", true},
				{"Expectancy is low", "C", false},
				{"Power needs dominate", "D", false},
			},
		},
		{
			prompt: `A PM introduces stretch assignments, public celebration of innovation, and increased autonomy in task selection.
Team members report higher pride in their outcomes even though salary did not change.
Which Herzberg factor primarily increased motivation?`,
			domain:          "Team Motivation Drill",
			explanation:     "Achievement, recognition, and responsibility are intrinsic motivators in Herzberg's two-factor theory, driving satisfaction.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Improved hygiene", "A", false},
				{"Enhanced motivators", "B", true},
				{"Higher safety needs", "C", false},
				{"Lower expectancy", "D", false},
			},
		},
		{
			prompt: `A functional manager monitors hours closely, limits decision-making authority, and believes employees avoid work unless supervised.
Team members request more autonomy but the manager fears performance would drop without tight control.
Which of McGregor's assumptions does this manager demonstrate?`,
			domain:          "Team Motivation Drill",
			explanation:     "Viewing people as disliking work and needing coercion reflects Theory X assumptions in McGregor's model.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Theory X", "A", true},
				{"Theory Y", "B", false},
				{"Theory Z", "C", false},
				{"Expectancy theory", "D", false},
			},
		},
		{
			prompt: `A self-managing agile squad proactively learns new tools, proposes experiments, and volunteers for complex work.
The coach spends minimal time directing tasks because members own their commitments.
Which motivational view aligns with this scenario?`,
			domain:          "Team Motivation Drill",
			explanation:     "Employees who seek responsibility and self-direction align with McGregor's Theory Y assumptions.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Theory X", "A", false},
				{"Theory Y", "B", true},
				{"Theory Z", "C", false},
				{"Expectancy theory", "D", false},
			},
		},
		{
			prompt: `A business analyst thrives on solving complex defects before anyone else and takes pride in receiving stretch tasks.
They are less interested in recognition events and more motivated by opportunities to conquer difficult problems.
According to McClelland's needs theory, which need is dominant?`,
			domain:          "Team Motivation Drill",
			explanation:     "A drive to master challenging tasks and excel individually signals a high need for Achievement (nAch).",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Need for Affiliation", "A", false},
				{"Need for Achievement", "B", true},
				{"Need for Power", "C", false},
				{"Need for Security", "D", false},
			},
		},
		{
			prompt: `A senior architect pushes to lead enterprise design authority, preferring assignments where they influence strategy and resource allocation.
They lobby for presentation slots with the CIO and enjoy steering debates.
Which McClelland need is most evident?`,
			domain:          "Team Motivation Drill",
			explanation:     "Seeking influence and control over direction reflects a strong Need for Power (nPow).",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Need for Affiliation", "A", false},
				{"Need for Power", "B", true},
				{"Need for Achievement", "C", false},
				{"Need for Security", "D", false},
			},
		},
		{
			prompt: `A change analyst prefers collaborative workshops, gravitates toward mentoring new hires, and worries about upsetting team cohesion.
They choose tasks that keep them connected to others even if technical complexity is moderate.
Which McClelland need is most likely dominant?`,
			domain:          "Team Motivation Drill",
			explanation:     "Prioritising relationships and harmony indicates a high Need for Affiliation (nAff).",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Need for Achievement", "A", false},
				{"Need for Power", "B", false},
				{"Need for Affiliation", "C", true},
				{"Need for Security", "D", false},
			},
		},
		{
			prompt: `After a bonus freeze, engineers say extra effort won't matter because the company never ties performance to tangible rewards.
They still believe they can meet technical targets, but doubt those results will change their compensation or recognition.
Which element of Vroom's expectancy theory is weakest?`,
			domain:          "Team Motivation Drill",
			explanation:     "If employees believe performance will not lead to rewards, instrumentality is low, undermining motivation.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Expectancy", "A", false},
				{"Instrumentality", "B", true},
				{"Valence", "C", false},
				{"Equity", "D", false},
			},
		},
		{
			prompt: `A product owner offers certification vouchers that team members highly value, and explains the specific performance goals required to earn them.
Developers believe the goals are attainable and trust the organisation will deliver the promised reward.
Which combination of Vroom factors is being reinforced?`,
			domain:          "Team Motivation Drill",
			explanation:     "Clear linkage between effort, successful performance, and a valued reward strengthens expectancy, instrumentality, and valence.",
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"Only expectancy", "A", false},
				{"Instrumentality and valence", "B", false},
				{"Expectancy, instrumentality, and valence together", "C", true},
				{"Only valence", "D", false},
			},
		},
		{
			prompt: `A manager introduces a recognition program but employees shrug because the reward is a parking space none of them use.
They understand how to earn it, yet the prize holds little personal appeal.
Which Vroom component is failing?`,
			domain:          "Team Motivation Drill",
			explanation:     "If a reward is not valued, valence is low, so motivation remains weak even when expectancy and instrumentality are intact.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Expectancy", "A", false},
				{"Instrumentality", "B", false},
				{"Valence", "C", true},
				{"Equity", "D", false},
			},
		},
		{
			prompt: `A Japanese-owned subsidiary fosters long-term employment, rotates staff across functions, and builds high trust between management and teams.
Decision-making emphasises consensus and loyalty, and employees feel the company invests in their whole career.
Which motivational theory best captures this philosophy?`,
			domain:          "Team Motivation Drill",
			explanation:     "Ouchi's Theory Z emphasises trust, holistic concern, and long-term employment to motivate employees.",
			popularityScore: 3.3,
			choices: []ChoiceData{
				{"Theory X", "A", false},
				{"Theory Y", "B", false},
				{"Theory Z", "C", true},
				{"Expectancy theory", "D", false},
			},
		},
		{
			prompt: `A project sponsor wants to reinvigorate a plateauing team.
They propose salary adjustments, upgraded laptops, and improved cafeteria options but offer no career growth or recognition opportunities.
What would Herzberg predict about motivational impact?`,
			domain:          "Team Motivation Drill",
			explanation:     "Improving hygiene factors addresses dissatisfaction but without motivators intrinsic motivation will remain flat.",
			popularityScore: 3.4,
			choices: []ChoiceData{
				{"Motivation will surge because hygiene factors were the missing ingredient", "A", false},
				{"Dissatisfaction may drop but true motivation will not significantly increase", "B", true},
				{"Employees will focus on safety needs", "C", false},
				{"Expectancy will collapse", "D", false},
			},
		},
		{
			prompt: `A team experiences a leadership change and suddenly debates every user story definition, creating friction.
Previously agreed working agreements are ignored and sub-groups form around different opinions.
The Scrum Master notes they had been collaborating smoothly for months before the disruption.
How should the team interpret this regression using Tuckman's model?`,
			domain:          "Team Motivation Drill",
			explanation:     "Significant change can push a previously normed team back into Storming as roles and power dynamics are renegotiated.",
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"They remain in Performing and should escalate dissent", "A", false},
				{"They have slipped into Storming and need facilitation to rebuild norms", "B", true},
				{"They moved directly to Adjourning", "C", false},
				{"They returned to Forming and must immediately rewrite the charter", "D", false},
			},
		},
		{
			prompt: `A domain expert says, "I trust any stretch assignment will lead to promotion, so I'm willing to take on the new portfolio dashboard."
Another teammate says, "Even if I work nights, leadership never notices, so why bother?"
Which motivational framework helps the PM address both perspectives?`,
			domain:          "Team Motivation Drill",
			explanation:     "Vroom's expectancy theory links effort, performance, and reward; differing beliefs about instrumentality explain the contrasting motivation levels.",
			popularityScore: 3.5,
			choices: []ChoiceData{
				{"Herzberg's hygiene theory", "A", false},
				{"Vroom's expectancy theory", "B", true},
				{"Theory Z", "C", false},
				{"McClelland's needs theory", "D", false},
			},
		},
		{
			prompt: `A remote analytics squad requests more time for virtual coffee chats, buddy systems for new hires, and shared recognition rituals.
They consistently meet sprint goals but say the lack of social glue drains their energy.
Which Maslow need is the team highlighting?`,
			domain:          "Team Motivation Drill",
			explanation:     "Desire for belonging, camaraderie, and social bonding reflects the Social layer of Maslow's hierarchy.",
			popularityScore: 3.2,
			choices: []ChoiceData{
				{"Physiological", "A", false},
				{"Safety", "B", false},
				{"Social", "C", true},
				{"Esteem", "D", false},
			},
		},
	}

	return questions
}

func round2(value float64) float64 {
	return math.Round(value*100) / 100
}

func formatCurrency(value float64) string {
	return fmt.Sprintf("RM%.2f", value)
}

func formatPercent(value float64) string {
	return fmt.Sprintf("%.0f%%", math.Round(value*100))
}

func formatRatio(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.005
}

func pertDrillQuestions() []QuestionData {
	const domain = "PERT Drill"

	questions := make([]QuestionData, 0, 50)
	rotation := 0

	teScenarios := []struct {
		name        string
		optimistic  float64
		mostLikely  float64
		pessimistic float64
	}{
		{"Atlas", 4, 6, 12},
		{"Beacon", 5, 8, 14},
		{"Comet", 3, 7, 11},
		{"Dynamo", 6, 9, 15},
		{"Eclipse", 2, 5, 9},
		{"Frontier", 8, 12, 18},
		{"Galaxy", 10, 14, 22},
		{"Harbor", 7, 10, 16},
		{"Inertia", 9, 13, 20},
		{"Jupiter", 11, 16, 26},
		{"Kepler", 5, 9, 13},
		{"Lighthouse", 4, 7, 12},
		{"Mirage", 6, 8, 11},
		{"Nimbus", 3, 6, 10},
		{"Orion", 2, 6, 9},
		{"Pinnacle", 1, 5, 8},
		{"Quasar", 3, 8, 12},
		{"Radiant", 4, 10, 18},
		{"Summit", 5, 11, 19},
		{"Titan", 6, 12, 21},
	}

	for i, s := range teScenarios {
		te := round2((s.optimistic + 4*s.mostLikely + s.pessimistic) / 6)
		distractors := numericDistractors(te)
		explanation := fmt.Sprintf("TE = (O + 4M + P) / 6 = (%.2f + 4×%.2f + %.2f) / 6 = %.2f.", s.optimistic, s.mostLikely, s.pessimistic, te)
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[PERT TE] Activity %s has O = %.2f days, M = %.2f days, and P = %.2f days. What is the expected duration (TE)?", s.name, s.optimistic, s.mostLikely, s.pessimistic),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.4,
			choices:         createChoices(te, distractors, rotation+i, formatNumber),
		})
	}

	sigmaScenarios := []struct {
		name        string
		optimistic  float64
		pessimistic float64
	}{
		{"Umbra", 4, 14},
		{"Vanguard", 6, 18},
		{"Waypoint", 2, 8},
		{"Xenon", 5, 11},
		{"Yukon", 3, 15},
		{"Zenith", 7, 13},
		{"Apex", 1, 7},
		{"Blaze", 2, 10},
		{"Cinder", 4, 16},
		{"Drift", 6, 12},
		{"Ember", 3, 9},
		{"Flare", 5, 17},
		{"Glide", 4, 10},
		{"Horizon", 8, 14},
		{"Inferno", 9, 21},
	}

	sigmaRotationBase := rotation + len(teScenarios)
	for i, s := range sigmaScenarios {
		sigma := round2((s.pessimistic - s.optimistic) / 6)
		distractors := numericDistractors(sigma)
		explanation := fmt.Sprintf("σ = (P - O) / 6 = (%.2f - %.2f) / 6 = %.2f.", s.pessimistic, s.optimistic, sigma)
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[PERT σ] Activity %s has O = %.2f days and P = %.2f days. What is the standard deviation?", s.name, s.optimistic, s.pessimistic),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.4,
			choices:         createChoices(sigma, distractors, sigmaRotationBase+i, formatNumber),
		})
	}

	varianceScenarios := []struct {
		name  string
		sigma float64
	}{
		{"Journey", 1.5},
		{"Keystone", 1.2},
		{"Lattice", 2.0},
		{"Matrix", 1.8},
		{"Nexus", 0.9},
		{"Orbit", 1.1},
		{"Paragon", 1.6},
		{"Quarry", 2.2},
		{"Resonance", 1.3},
		{"Stratus", 1.7},
		{"Tundra", 0.8},
		{"Utopia", 1.4},
		{"Velocity", 1.0},
		{"Whisper", 1.9},
		{"Zephyr", 2.1},
	}

	varianceRotationBase := sigmaRotationBase + len(sigmaScenarios)
	for i, s := range varianceScenarios {
		variance := round2(s.sigma * s.sigma)
		distractors := numericDistractors(variance)
		explanation := fmt.Sprintf("Variance = σ² = %.2f² = %.2f.", s.sigma, variance)
		questions = append(questions, QuestionData{
			prompt:          fmt.Sprintf("[PERT Variance] Activity %s has σ = %.2f. What is the variance?", s.name, s.sigma),
			domain:          domain,
			explanation:     explanation,
			popularityScore: 3.3,
			choices:         createChoices(variance, distractors, varianceRotationBase+i, formatNumber),
		})
	}

	return questions
}

func numericDistractors(correct float64) []float64 {
	adjustments := []float64{-0.2, -0.12, 0.1, 0.18, -0.08, 0.15}
	values := make([]float64, 0, 3)
	for _, adj := range adjustments {
		candidate := round2(correct * (1 + adj))
		if candidate <= 0 && correct > 0 {
			candidate = round2(correct * (1 - adj))
		}
		if almostEqual(candidate, correct) {
			continue
		}
		duplicate := false
		for _, existing := range values {
			if almostEqual(existing, candidate) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			values = append(values, candidate)
		}
		if len(values) == 3 {
			break
		}
	}
	if len(values) < 3 {
		fallbacks := []float64{round2(correct + 0.45), round2(correct - 0.38), round2(correct + 0.72)}
		for _, fb := range fallbacks {
			if almostEqual(fb, correct) {
				continue
			}
			duplicate := false
			for _, existing := range values {
				if almostEqual(existing, fb) {
					duplicate = true
					break
				}
			}
			if !duplicate {
				values = append(values, fb)
			}
			if len(values) == 3 {
				break
			}
		}
	}
	return values
}

func currencyDistractors(correct float64) []float64 {
	adjustments := []float64{-0.18, -0.1, 0.12, 0.2, -0.05, 0.08}
	values := make([]float64, 0, 3)
	for _, adj := range adjustments {
		candidate := round2(correct * (1 + adj))
		if candidate <= 0 && correct > 0 {
			candidate = round2(correct * (1 - adj))
		}
		if almostEqual(candidate, correct) {
			continue
		}
		duplicate := false
		for _, existing := range values {
			if almostEqual(existing, candidate) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			values = append(values, candidate)
		}
		if len(values) == 3 {
			break
		}
	}
	if len(values) < 3 {
		fallbacks := []float64{round2(correct + 5000), round2(correct - 4000), round2(correct + 2500)}
		for _, fb := range fallbacks {
			if almostEqual(fb, correct) {
				continue
			}
			duplicate := false
			for _, existing := range values {
				if almostEqual(existing, fb) {
					duplicate = true
					break
				}
			}
			if !duplicate {
				values = append(values, fb)
			}
			if len(values) == 3 {
				break
			}
		}
	}
	return values
}

func ratioDistractors(correct float64) []float64 {
	adjustments := []float64{-0.15, -0.08, 0.1, 0.18, -0.05, 0.06}
	values := make([]float64, 0, 3)
	for _, adj := range adjustments {
		candidate := round2(correct * (1 + adj))
		if candidate <= 0 && correct > 0 {
			candidate = round2(correct * (1 - adj))
		}
		if almostEqual(candidate, correct) {
			continue
		}
		duplicate := false
		for _, existing := range values {
			if almostEqual(existing, candidate) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			values = append(values, candidate)
		}
		if len(values) == 3 {
			break
		}
	}
	if len(values) < 3 {
		fallbacks := []float64{round2(correct + 0.12), round2(correct - 0.09), round2(correct + 0.18)}
		for _, fb := range fallbacks {
			if almostEqual(fb, correct) {
				continue
			}
			duplicate := false
			for _, existing := range values {
				if almostEqual(existing, fb) {
					duplicate = true
					break
				}
			}
			if !duplicate {
				values = append(values, fb)
			}
			if len(values) == 3 {
				break
			}
		}
	}
	return values
}

func formatNumber(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func createChoices(correct float64, distractors []float64, rotation int, formatter func(float64) string) []ChoiceData {
	values := append([]float64{correct}, distractors...)
	if len(values) < 4 {
		for len(values) < 4 {
			increment := 500.0
			absVal := math.Abs(correct)
			if absVal <= 50 {
				increment = 5
			}
			if absVal <= 10 {
				increment = 1
			}
			if absVal <= 1 {
				increment = 0.2
			}
			values = append(values, round2(correct+float64(len(values))*increment))
		}
	}
	order := []int{0, 1, 2, 3}
	rotation = rotation % len(order)
	reordered := append(order[rotation:], order[:rotation]...)
	choices := make([]ChoiceData, len(order))
	for i, idx := range reordered {
		value := values[idx]
		choices[i] = ChoiceData{
			text:      formatter(value),
			label:     string(rune('A' + i)),
			isCorrect: idx == 0,
		}
	}
	return choices
}

// GetAdditionalQuestions returns supplementary questions (currently none)
func GetAdditionalQuestions() []QuestionData {
	return []QuestionData{}
}
