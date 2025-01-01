package modular

func AllBlocks() []BlockInterface {
	return []BlockInterface{
		&MarketingHeroCoverImageWithCtas{
			Type:        "MarketingHeroCoverImageWithCtas",
			HideFromNav: true,
			Heading:     "Discover new product and best possibilities",
			Subheading:  "Here at Flowbite we focus on markets where technology, innovation, and capital can unlock long-term value and drive economic growth.",
			Left: Side{
				Heading:    "28 November 2021",
				Subheading: "Join us at FlowBite 2021 to understand what's next as the global tech and startup ecosystem, rethinks the future of everything.",
				CTA: CTA{
					Text: "Conference",
					URL:  "#",
				},
			},
			Right: Side{
				Heading:    "25+ top notch speakers",
				Subheading: "Here you will find keynote speakers, who all are able to talk about Recruiting. Click on the individual keynote speakers and read more about them and their keynotes.",
				CTA: CTA{
					Text: "View list",
					URL:  "#",
				},
			},
			Image: Image{
				URL: "https://flowbite.s3.amazonaws.com/blocks/marketing-ui/hero/conference-speaker.jpg",
			},
		},
		&FeatureSection{
			Type:        "FeatureSection",
			Library:     "FlowBite",
			View:        "cta-list",
			Heading:     "The most trusted cryptocurrency platform",
			HideFromNav: false,
			Subheading:  "Here are a few reasons why you should choose Flowbite",
			Features: []Feature{
				{
					Heading: "Secure storage",
					Summary: "We store the vast majority of the digital assets in secure offline storage.",
					Icon:    "solid-wand-magic-sparkles",
				},
				{
					Heading: "Insurance",
					Summary: "Flowbite maintains crypto insurance and all USD cash balances are covered.",
					Icon:    "solid-award",
				},
				{
					Heading: "Best Practices",
					Summary: "Flowbite marketplace supports a variety of the most popular digital currencies.",
					Icon:    "solid-badge-check",
				},
			},
		},
		&FeatureSection{
			Type:        "FeatureSection",
			Library:     "FlowBite",
			Heading:     "Designed for business teams like yours",
			HideFromNav: false,
			Subheading:  "Here are a few reasons why you should choose Flowbite",
			View:        "icons",
			Features: []Feature{
				{
					Heading: "Just the right balance for growth",
					Summary: "Enterprise tools cost more, are difficult to use, and take longer to implement. According to G2, Flowbite is the easiest-to-use tool, with the fastest time to ROI.",
					Icon:    "solid-clock",
				},
				{
					Heading: "Just the right balance for growth",
					Summary: "Enterprise tools cost more, are difficult to use, and take longer to implement. According to G2, Flowbite is the easiest-to-use tool, with the fastest time to ROI.",
					Icon:    "solid-heart",
				},
			},
		},
		&FeatureSection{
			Type:        "FeatureSection",
			Library:     "FlowBite",
			HideFromNav: false,
			Heading:     "Secure platform, secure data",
			Subheading:  "Here are a few reasons why you should choose Flowbite",
			View:        "card-list",
			Features: []Feature{
				{
					Heading: "Marketing",
					Summary: "Plan it, create it, launch it. Collaborate seamlessly with all the organization and hit your marketing goals every month with our marketing plan.",
				},
				{
					Heading: "Legal",
					Summary: "Protect your organization, devices and stay compliant with our structured workflows and custom permissions made for you.",
				},
				{
					Heading: "Business Automation",
					Summary: "Auto-assign tasks, send Slack messages, and much more. Now power up with hundreds of new templates to help you get started.",
				},
				{
					Heading: "Finance",
					Summary: "Audit-proof software built for critical financial operations like month-end close and quarterly budgeting.",
				},
				{
					Heading: "Enterprise Design",
					Summary: "Craft beautiful, delightful experiences for both marketing and product with real cross-company collaboration.",
				},
				{
					Heading: "Operations",
					Summary: "Keep your company's lights on with customizable, iterative, and structured workflows built for all efficient teams and individual.",
				},
			},
		},
		&PricingTable{
			Type:          "PricingTable",
			HideFromNav:   false,
			Library:       "FlowBite",
			Heading:       "Pricing Plans",
			Subheading:    "Flexible pricing for all teams and budgets.",
			ProductLineID: "observability_workshop_1",
		},
		&ContactFormsDefault{
			Type:        "ContactFormsDefault",
			HideFromNav: false,
			Library:     "FlowBite",
			Heading:     "Contact us",
			Subheading:  "We are a team of developers and designers who are passionate about creating beautiful and functional websites.",
		},
		&HeroSectionsDefault{
			Type:        "HeroSectionsDefault",
			HideFromNav: false,
			Library:     "FlowBite",
			Heading:     "Hero Section",
			Subheading:  "We are a team of developers and designers who are passionate about creating beautiful and functional websites.",
			Left: SideJustCTA{
				CTA: CTA{
					Text: "Learn More Today",
					URL:  "#",
				},
			},
			Right: SideJustCTA{
				CTA: CTA{
					Text: "Learn More Today",
					URL:  "#",
				},
			},
		},
		&TestimonialSectionsBlockquote{
			Type:        "TestimonialSectionsBlockquote",
			HideFromNav: false,
			Library:     "FlowBite",
			Testimonial: Testimonial{
				Content: "We are a team of developers and designers who are passionate about creating beautiful and functional websites.",
				Author: Author{
					Name:     "Micheal Gough",
					Title:    "CEO at Google",
					ImageSrc: "https://flowbite.s3.amazonaws.com/blocks/marketing-ui/avatars/michael-gouch.png",
				},
			},
		},
		&SocialProofCardStatistics{
			Type:        "SocialProofCardStatistics",
			HideFromNav: false,
			Library:     "FlowBite",
			Heading:     "Social Proof Card Statistics",
			Subheading:  "We are a team of developers and designers who are passionate about creating beautiful and functional websites.",
			Note:        "Results based on a composite organization of 1,800 developers using GitHub over three years.",
			Sentiment:   "neutral",
			Statistics: Statistics{
				Left: Statistic{
					Value:       "40%",
					Title:       "Reduction",
					Description: "in developer onboarding time",
				},
				Center: Statistic{
					Value:       "469%",
					Title:       "Return on investment",
					Description: "over 3 years",
				},
				Right: Statistic{
					Value:       "60+",
					Title:       "Minutes saved",
					Description: "per developer, per day",
				},
			},
		},
		&FeatureSectionsAlternate{
			Type:        "FeatureSectionsAlternate",
			HideFromNav: false,
			Library:     "FlowBite",
			Left: BulletList{
				Heading:    "Feature Sections Alternate",
				Subheading: "Feature Sections Alternate",
				Footer:     "Footer",
				Sentiment:  "negative",
				Feature1:   "Feature 1",
				Feature2:   "Feature 2",
				Feature3:   "Feature 3",
				Feature4:   "Feature 4",
				Feature5:   "Feature 5",
				Feature6:   "Feature 6",
			},
			Right: BulletList{
				Heading:    "Feature Sections Alternate",
				Subheading: "Feature Sections Alternate",
				Footer:     "Footer",
				Sentiment:  "positive",
				Feature1:   "Feature 1",
				Feature2:   "Feature 2",
				Feature3:   "Feature 3",
				Feature4:   "Feature 4",
				Feature5:   "Feature 5",
				Feature6:   "Feature 6",
			},
		},
	}
}
