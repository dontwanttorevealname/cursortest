package templates

import (
    "math"
    "math/rand"
    "sort"
)

var (
    commonTechPost = Post{
        Title:       "The Future of AI Development",
        Description: "Discussing the latest breakthroughs in artificial intelligence and what they mean for developers.",
        Comments:    156,
        Likes:       432,
        PondName:    "TechTalk",
        Author:      "ByteMaster",
    }

    commonGardenPost = Post{
        Title:       "Spring Gardening Tips",
        Description: "Essential tips for preparing your garden for the spring season. Share your experiences!",
        Comments:    89,
        Likes:       245,
        PondName:    "GreenThumb",
        Author:      "GreenGuru",
    }

    commonArtPost = Post{
        Title:       "Digital Art Techniques Workshop",
        Description: "Join our weekly digital art workshop! This week we're focusing on character design.",
        Comments:    167,
        Likes:       389,
        PondName:    "ArtistsCorner",
        Author:      "PixelPainter",
    }

    commonBookPost = Post{
        Title:       "Monthly Book Discussion: Sci-Fi Classics",
        Description: "Join us as we explore the foundational works of science fiction literature.",
        Comments:    142,
        Likes:       276,
        PondName:    "BookClub",
        Author:      "BookWorm42",
    }

    commonFoodPost = Post{
        Title:       "Seasonal Cooking with Home-Grown Herbs",
        Description: "Make the most of your garden's bounty with these delicious recipes!",
        Comments:    98,
        Likes:       312,
        PondName:    "FoodiesUnite",
        Author:      "ChefCroak",
    }

    commonSciencePost = Post{
        Title:       "Latest Discoveries in Quantum Physics",
        Description: "Breaking down the newest research in quantum mechanics and its implications.",
        Comments:    187,
        Likes:       423,
        PondName:    "ScienceLab",
        Author:      "LabRat",
    }

    // Posts from users without accounts
    randomTechPost = Post{
        Title:       "Building a Home Server Setup",
        Description: "My journey setting up a home media and development server. Here's what I learned.",
        Comments:    92,
        Likes:       278,
        PondName:    "TechTalk",
        Author:      "ServerMaster",
    }

    randomGardenPost = Post{
        Title:       "Desert Plants Survival Guide",
        Description: "Tips for maintaining a thriving garden in arid climates.",
        Comments:    67,
        Likes:       189,
        PondName:    "GreenThumb",
        Author:      "DesertBloom",
    }

    randomArtPost = Post{
        Title:       "Traditional vs Digital Art",
        Description: "Breaking down the pros and cons of both mediums from years of experience.",
        Comments:    156,
        Likes:       445,
        PondName:    "ArtistsCorner",
        Author:      "ArtisanSage",
    }

    randomFoodPost = Post{
        Title:       "International Street Food Series",
        Description: "Exploring street food cultures around the world. First stop: Thailand!",
        Comments:    134,
        Likes:       367,
        PondName:    "FoodiesUnite",
        Author:      "StreetFoodie",
    }

    // Update official posts to include author display names
    officialWelcomePost = Post{
        Title:       "Welcome to Ribbit!",
        Description: "Get started with your journey here. Learn about our community guidelines and features.",
        Comments:    42,
        Likes:       156,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    }

    officialUpdatePost = Post{
        Title:       "New Features: Enhanced Pond Navigation",
        Description: "We've updated how you browse and discover new ponds. Check out the new features!",
        Comments:    83,
        Likes:       245,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    }

    officialCommunityPost = Post{
        Title:       "Community Spotlight: February Stars",
        Description: "Celebrating our most active contributors and their amazing content this month.",
        Comments:    67,
        Likes:       198,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    }

    officialGuidelinesPost = Post{
        Title:       "Updated Community Guidelines",
        Description: "Important updates to our community guidelines. Please review these changes.",
        Comments:    91,
        Likes:       267,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    }


)

type UserTemplate struct {
    Email         string
    Password      string
    OfficialPosts []Post
    Posts         []Post
    Ponds         []Pond
}

type Post struct {
    Title       string
    Description string
    Comments    int
    Likes       int
    PondName    string
    Author      string    // Add author field
    Type        string    // Add type field
}

type Pond struct {
    Name        string
    Description string
    Members     string
}

// Helper function to get random posts from a pond
func getRandomPostsFromPond(posts []Post, minCount int, excludeAuthor string) []Post {
    available := make([]Post, 0)
    for _, post := range posts {
        if post.Author != excludeAuthor {
            available = append(available, post)
        }
    }

    if len(available) == 0 {
        return []Post{}
    }

    // Shuffle the posts
    rand.Shuffle(len(available), func(i, j int) {
        available[i], available[j] = available[j], available[i]
    })

    // Return random number of posts between minCount and minCount+3
    count := minCount + rand.Intn(4)
    if count > len(available) {
        return available
    }
    return available[:count]
}

// Get random posts for a user based on their pond memberships
func getRandomPostsForUser(userPonds []Pond, excludeAuthor string) []Post {
    var allPosts []Post
    
    // Get random number of posts from each pond
    for _, pond := range userPonds {
        var pondPosts []Post
        switch pond.Name {
        case "TechTalk":
            pondPosts = getRandomPostsFromPond(techPosts, 3, excludeAuthor)
        case "GreenThumb":
            pondPosts = getRandomPostsFromPond(gardenPosts, 3, excludeAuthor)
        case "ArtistsCorner":
            pondPosts = getRandomPostsFromPond(artPosts, 3, excludeAuthor)
        case "FoodiesUnite":
            pondPosts = getRandomPostsFromPond(foodPosts, 3, excludeAuthor)
        case "ScienceLab":
            pondPosts = getRandomPostsFromPond(sciencePosts, 3, excludeAuthor)
        case "BookClub":
            pondPosts = getRandomPostsFromPond(bookPosts, 3, excludeAuthor)
        case "CodingPond":
            pondPosts = getRandomPostsFromPond(codingPosts, 3, excludeAuthor)
        }
        allPosts = append(allPosts, pondPosts...)
    }

    // Shuffle all selected posts
    rand.Shuffle(len(allPosts), func(i, j int) {
        allPosts[i], allPosts[j] = allPosts[j], allPosts[i]
    })

    // Return random number between 6 and 10 posts
    numPosts := 6 + rand.Intn(5)
    if numPosts > len(allPosts) {
        return allPosts
    }
    return allPosts[:numPosts]
}

// GetUserTemplate returns the template for a specific user
func GetUserTemplate(username string) *UserTemplate {
    switch username {
    case "admin@ribbit.com":
        return getAdminTemplate()
    case "ByteMaster":
        return getTechUserTemplate()
    case "GreenGuru":
        return getGardenUserTemplate()
    case "PixelPainter":
        return getArtistTemplate()
    case "BookWorm42":
        return getBookLoverTemplate()
    case "ChefCroak":
        return getChefTemplate()
    case "LabRat":
        return getScientistTemplate()
    }
    return nil
}

// Create a consistent set of official posts that all users will see
var officialPosts = []Post{
    {
        Title:       "Welcome to Ribbit!",
        Description: "Get started with your journey here. Learn about our community guidelines and features.",
        Comments:    42,
        Likes:       156,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    },
    {
        Title:       "New Features: Enhanced Pond Navigation",
        Description: "We've updated how you browse and discover new ponds. Check out the new features!",
        Comments:    83,
        Likes:       245,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    },
    {
        Title:       "Community Spotlight: February Stars",
        Description: "Celebrating our most active contributors and their amazing content this month.",
        Comments:    67,
        Likes:       198,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    },
    {
        Title:       "Updated Community Guidelines",
        Description: "Important updates to our community guidelines. Please review these changes.",
        Comments:    91,
        Likes:       267,
        PondName:    "Official",
        Author:      "Ribbit Admin",
    },
}

// Update getAdminTemplate to include random posts
func getAdminTemplate() *UserTemplate {
    allPonds := []Pond{
        {Name: "TechTalk", Description: "All things technology", Members: "15K"},
        {Name: "GreenThumb", Description: "Gardening enthusiasts", Members: "8.5K"},
        {Name: "CodingPond", Description: "Programming discussions", Members: "12K"},
        {Name: "ArtistsCorner", Description: "Share your creations", Members: "9.8K"},
        {Name: "ScienceLab", Description: "Scientific discoveries and discussions", Members: "11.2K"},
        {Name: "BookClub", Description: "For literature lovers", Members: "7.3K"},
        {Name: "FoodiesUnite", Description: "Cooking and recipes", Members: "13.4K"},
    }

    return &UserTemplate{
        Email:         "admin@ribbit.com",
        Password:      "RibbitAdmin",
        OfficialPosts: officialPosts,  // Only actual official posts
        Posts:         getRandomPostsForUser(allPonds, "admin@ribbit.com"),
        Ponds:         allPonds,
    }
}

func getTechUserTemplate() *UserTemplate {
    userPonds := []Pond{
        {Name: "TechTalk", Description: "All things technology", Members: "15K"},
        {Name: "CodingPond", Description: "Programming discussions", Members: "12K"},
        {Name: "ScienceLab", Description: "Scientific discoveries and discussions", Members: "11.2K"},
    }

    return &UserTemplate{
        Email:         "ByteMaster",
        Password:      "Techn0l0gy!",
        OfficialPosts: officialPosts,  // Use consistent official posts
        Posts:         getRandomPostsForUser(userPonds, "ByteMaster"),
        Ponds:         userPonds,
    }
}

func getGardenUserTemplate() *UserTemplate {
    userPonds := []Pond{
        {Name: "GreenThumb", Description: "Gardening enthusiasts", Members: "8.5K"},
        {Name: "ScienceLab", Description: "Scientific discoveries and discussions", Members: "11.2K"},
    }

    return &UserTemplate{
        Email:         "GreenGuru",
        Password:      "Plant123!",
        OfficialPosts: officialPosts,
        Posts:         getRandomPostsForUser(userPonds, "GreenGuru"), // Will only get posts from GreenThumb and ScienceLab
        Ponds:         userPonds,
    }
}

func getArtistTemplate() *UserTemplate {
    userPonds := []Pond{
        {Name: "ArtistsCorner", Description: "Share your creations", Members: "9.8K"},
        {Name: "TechTalk", Description: "All things technology", Members: "15K"},
    }

    return &UserTemplate{
        Email:         "PixelPainter",
        Password:      "Art1st!",
        OfficialPosts: officialPosts,
        Posts:         getRandomPostsForUser(userPonds, "PixelPainter"),
        Ponds:         userPonds,
    }
}

func getBookLoverTemplate() *UserTemplate {
    userPonds := []Pond{
        {Name: "BookClub", Description: "For literature lovers", Members: "7.3K"},
        {Name: "ArtistsCorner", Description: "Share your creations", Members: "9.8K"},
    }

    return &UserTemplate{
        Email:         "BookWorm42",
        Password:      "ReadMore!",
        OfficialPosts: officialPosts,
        Posts:         getRandomPostsForUser(userPonds, "BookWorm42"),
        Ponds:         userPonds,
    }
}

func getChefTemplate() *UserTemplate {
    userPonds := []Pond{
        {Name: "FoodiesUnite", Description: "Cooking and recipes", Members: "13.4K"},
        {Name: "GreenThumb", Description: "Gardening enthusiasts", Members: "8.5K"},
    }

    return &UserTemplate{
        Email:         "ChefCroak",
        Password:      "Food1e!",
        OfficialPosts: officialPosts,
        Posts:         getRandomPostsForUser(userPonds, "ChefCroak"),
        Ponds:         userPonds,
    }
}

func getScientistTemplate() *UserTemplate {
    userPonds := []Pond{
        {Name: "ScienceLab", Description: "Scientific discoveries and discussions", Members: "11.2K"},
        {Name: "TechTalk", Description: "All things technology", Members: "15K"},
        {Name: "BookClub", Description: "For literature lovers", Members: "7.3K"},
    }

    return &UserTemplate{
        Email:         "LabRat",
        Password:      "Sc1ence!",
        OfficialPosts: officialPosts,
        Posts:         getRandomPostsForUser(userPonds, "LabRat"),
        Ponds:         userPonds,
    }
}

// GetTrendingPosts returns the top post from each pond, sorted by engagement
func GetTrendingPosts() []Post {
    pondBestPosts := make(map[string]Post)
    
    // Helper function to calculate engagement score
    getEngagementScore := func(post Post) float64 {
        // Calculate ratio of comments to likes
        // Add 1 to prevent division by zero and to smooth out the ratio
        commentRatio := float64(post.Comments+1) / float64(post.Likes+1)
        
        // Weight the score to favor posts with more total engagement
        totalEngagement := float64(post.Comments + post.Likes)
        
        // Combine ratio and total engagement for final score
        // This formula gives higher weight to posts with good comment/like ratio
        // while still considering overall engagement
        return commentRatio * math.Log(totalEngagement+1)
    }
    
    // Process all post collections
    processPostList := func(posts []Post) {
        for _, post := range posts {
            currentBest, exists := pondBestPosts[post.PondName]
            if !exists || getEngagementScore(post) > getEngagementScore(currentBest) {
                pondBestPosts[post.PondName] = post
            }
        }
    }
    
    // Process all post collections
    processPostList(techPosts)
    processPostList(gardenPosts)
    processPostList(artPosts)
    processPostList(foodPosts)
    processPostList(sciencePosts)
    processPostList(bookPosts)
    processPostList(codingPosts)
    
    // Convert map to slice and sort by engagement score
    var trendingPosts []Post
    for _, post := range pondBestPosts {
        trendingPosts = append(trendingPosts, post)
    }
    
    // Sort posts by engagement score
    sort.Slice(trendingPosts, func(i, j int) bool {
        return getEngagementScore(trendingPosts[i]) > getEngagementScore(trendingPosts[j])
    })
    
    return trendingPosts
}

// Update the existing artPosts array to include all posts
var artPosts = []Post{
    commonArtPost,
    randomArtPost,
    {
        Title:       "Drew this at 3am instead of sleeping",
        Description: "Just had to get this character design out of my head. Still not 100% happy with the lighting but...",
        Comments:    234,
        Likes:       876,
        PondName:    "ArtistsCorner",
        Author:      "NightOwlArtist",
    },
    {
        Title:       "Can we talk about art theft on Twitter?",
        Description: "Found my work being sold as NFTs without permission. Here's what I did to get them taken down...",
        Comments:    445,
        Likes:       1234,
        PondName:    "ArtistsCorner",
        Author:      "PixelPainter",
    },
    {
        Title:       "Trying oils after 10 years of digital",
        Description: "Ok wow this is HARD. Mad respect to traditional artists. Here's my first attempt (be gentle)...",
        Comments:    122,
        Likes:       445,
        PondName:    "ArtistsCorner",
        Author:      "DigitalDabbler",
    },
    {
        Title:       "Mastering Digital Illustration",
        Description: "A comprehensive guide to creating professional digital illustrations using industry-standard tools.",
        Comments:    234,
        Likes:       567,
        PondName:    "ArtistsCorner",
        Author:      "DigitalDaVinci",
    },
    {
        Title:       "From Sketch to Masterpiece",
        Description: "My journey through a 30-day art challenge. See the progression and lessons learned.",
        Comments:    189,
        Likes:       445,
        PondName:    "ArtistsCorner",
        Author:      "SketchMaster",
    },
    {
        Title:       "Color Theory Deep Dive",
        Description: "Understanding color relationships and how to create harmonious palettes in your artwork.",
        Comments:    167,
        Likes:       389,
        PondName:    "ArtistsCorner",
        Author:      "ColorWhisperer",
    },
    {
        Title:       "3D Modeling for Beginners",
        Description: "Getting started with Blender - tips and tricks for creating your first 3D masterpiece.",
        Comments:    145,
        Likes:       298,
        PondName:    "ArtistsCorner",
        Author:      "3DArtist",
    },
    {
        Title:       "Animation Fundamentals",
        Description: "The 12 principles of animation and how to apply them in your digital work.",
        Comments:    178,
        Likes:       412,
        PondName:    "ArtistsCorner",
        Author:      "AnimationPro",
    },
}

// Update the existing bookPosts array
var bookPosts = []Post{
    commonBookPost,
    {
        Title:       "Classic Literature Analysis",
        Description: "Deep diving into the themes of Jane Austen's works and their modern relevance.",
        Comments:    156,
        Likes:       342,
        PondName:    "BookClub",
        Author:      "LitScholar",
    },
	{
		Title:       "Just finished Project Hail Mary and I'm EMOTIONAL",
		Description: "No spoilers but that ending?? I need to talk about this with someone who understands...",
		Comments:    234,
		Likes:       678,
		PondName:    "BookClub",
		Author:      "BookWorm42",
	},
	{
		Title:       "Unpopular opinion: physical books > e-readers",
		Description: "Yes they take up space. Yes they're harder to travel with. But there's just something about the smell...",
		Comments:    456,
		Likes:       789,
		PondName:    "BookClub",
		Author:      "PageTurner",
	},
	{
		Title:       "Building my reading nook! (progress pics)",
		Description: "Finally converted that weird corner of my apartment into the cozy reading space of my dreams",
		Comments:    123,
		Likes:       445,
		PondName:    "BookClub",
		Author:      "CozyReader",
	},
    {
        Title:       "Modern Fantasy Recommendations",
        Description: "A curated list of must-read fantasy novels from the last decade.",
        Comments:    189,
        Likes:       423,
        PondName:    "BookClub",
        Author:      "FantasyReader",
    },
    {
        Title:       "Writing Workshop: Character Development",
        Description: "Tips and exercises for creating memorable, three-dimensional characters.",
        Comments:    167,
        Likes:       356,
        PondName:    "BookClub",
        Author:      "WordSmith",
    },
    {
        Title:       "Book Photography Tips",
        Description: "How to take Instagram-worthy photos of your current reads.",
        Comments:    134,
        Likes:       445,
        PondName:    "BookClub",
        Author:      "BookishPhotographer",
    },
    {
        Title:       "Reading Challenge 2024",
        Description: "Join our annual reading challenge! 52 books, 52 weeks, endless adventures.",
        Comments:    223,
        Likes:       567,
        PondName:    "BookClub",
        Author:      "ReadingChampion",
    },
}

// Update the existing foodPosts array
var foodPosts = []Post{
    commonFoodPost,
    randomFoodPost,
    {
        Title:       "Baking Science Explained",
        Description: "Understanding the chemistry behind perfect pastries and breads.",
        Comments:    178,
        Likes:       389,
        PondName:    "FoodiesUnite",
        Author:      "Bakeologist",
    },
	{
		Title:       "Made grandma's secret pasta sauce (pics inside!)",
		Description: "Took me 3 tries but I finally got it exactly how she used to make it. The secret ingredient was...",
		Comments:    234,
		Likes:       567,
		PondName:    "FoodiesUnite",
		Author:      "PastaPro",
	},
	{
		Title:       "Hot take: Air fryers are overrated",
		Description: "Fight me in the comments but it's just a tiny convection oven and I'm tired of pretending it's not",
		Comments:    567,
		Likes:       890,
		PondName:    "FoodiesUnite",
		Author:      "ChefCroak",
	},
	{
		Title:       "I messed up my sourdough starter :(",
		Description: "Went on vacation and forgot to ask someone to feed Timothy (my starter). Any tips for starting over?",
		Comments:    89,
		Likes:       234,
		PondName:    "FoodiesUnite",
		Author:      "SourdoughSam",
	},
    {
        Title:       "Zero-Waste Cooking Guide",
        Description: "Creative ways to use every part of your ingredients and reduce kitchen waste.",
        Comments:    145,
        Likes:       334,
        PondName:    "FoodiesUnite",
        Author:      "SustainableChef",
    },
    {
        Title:       "World Cuisine Series: Thai",
        Description: "Essential ingredients and techniques for authentic Thai cooking at home.",
        Comments:    198,
        Likes:       445,
        PondName:    "FoodiesUnite",
        Author:      "GlobalGourmet",
    },
    {
        Title:       "Food Photography Tips",
        Description: "Professional tips for making your dishes look as good as they taste.",
        Comments:    167,
        Likes:       378,
        PondName:    "FoodiesUnite",
        Author:      "FoodLenser",
    },
}

// Update the existing gardenPosts array
var gardenPosts = []Post{
    commonGardenPost,
    randomGardenPost,
    {
        Title:       "Urban Balcony Gardens",
        Description: "Maximizing your small space for a thriving garden oasis.",
        Comments:    145,
        Likes:       312,
        PondName:    "GreenThumb",
        Author:      "UrbanGardener",
    },
	{
		Title:       "HELP! What's eating my tomatoes??",
		Description: "There are these weird marks on the leaves and something's definitely munching them. Pictures attached...",
		Comments:    45,
		Likes:       123,
		PondName:    "GreenThumb",
		Author:      "PlantPanic",
	},
	{
		Title:       "My monstera is trying to take over my apartment",
		Description: "Not complaining but this thing has grown 3 feet in 2 months?! Anyone else's monstera going crazy this summer?",
		Comments:    89,
		Likes:       234,
		PondName:    "GreenThumb",
		Author:      "MonsteraMan",
	},
	{
		Title:       "Secret weapon for pest control: LADYBUGS!",
		Description: "Y'all. I released 1500 ladybugs in my garden last week and the aphids are GONE. Best $20 I ever spent.",
		Comments:    167,
		Likes:       556,
		PondName:    "GreenThumb",
		Author:      "GreenGuru",
	},
	{
		Title:       "My vertical garden setup - 1 year later",
		Description: "Update on my balcony garden experiment. The strawberries are thriving!",
		Comments:    167,
		Likes:       445,
		PondName:    "GreenThumb",
		Author:      "UrbanGardener",
	},
	{
		Title:       "Winter gardening tips that actually work",
		Description: "Tested these methods in zone 5b. Even got fresh herbs in January!",
		Comments:    234,
		Likes:       567,
		PondName:    "GreenThumb",
		Author:      "FrostFarmer",
	},
	{
		Title:       "Help! Fungus gnats everywhere",
		Description: "They're taking over my indoor plants. Already tried neem oil...",
		Comments:    89,
		Likes:       234,
		PondName:    "GreenThumb",
		Author:      "PlantNewbie",
	},
	{
		Title:       "My composting journey (with pics)",
		Description: "From kitchen scraps to black gold. Here's what I learned...",
		Comments:    145,
		Likes:       389,
		PondName:    "GreenThumb",
		Author:      "CompostQueen",
	},
	{
		Title:       "Propagation success stories",
		Description: "Finally got a variegated monstera leaf to root! Method in comments.",
		Comments:    198,
		Likes:       567,
		PondName:    "GreenThumb",
		Author:      "PropagationStation",
	},

    {
        Title:       "Composting Masterclass",
        Description: "Everything you need to know about creating rich, healthy compost.",
        Comments:    167,
        Likes:       345,
        PondName:    "GreenThumb",
        Author:      "CompostKing",
    },
    {
        Title:       "Native Plant Guide",
        Description: "Why and how to incorporate native species into your garden.",
        Comments:    189,
        Likes:       423,
        PondName:    "GreenThumb",
        Author:      "NativePlantPro",
    },
    {
        Title:       "Indoor Plant Care 101",
        Description: "Essential tips for keeping your houseplants happy and healthy.",
        Comments:    156,
        Likes:       367,
        PondName:    "GreenThumb",
        Author:      "PlantParent",
    },
}

// Update the existing sciencePosts array
var sciencePosts = []Post{
    commonSciencePost,
    {
        Title:       "Space Exploration Updates",
        Description: "Latest developments in space technology and upcoming missions.",
        Comments:    234,
        Likes:       567,
        PondName:    "ScienceLab",
        Author:      "SpaceExplorer",
    },
	{
		Title:       "ELI5: The new quantum computing breakthrough",
		Description: "Trying to understand the IBM announcement. Can someone break down what this means for the field?",
		Comments:    234,
		Likes:       567,
		PondName:    "ScienceLab",
		Author:      "QuantumCurious",
	},
	{
		Title:       "Just published my first paper!!!",
		Description: "After 2 years of research, my work on neural network optimization is finally out! Link in comments",
		Comments:    345,
		Likes:       789,
		PondName:    "ScienceLab",
		Author:      "LabRat",
	},
	{
		Title:       "Cool physics demos for high school students?",
		Description: "Need to make physics exciting for my 10th graders. What demonstrations blew your mind as a student?",
		Comments:    123,
		Likes:       345,
		PondName:    "ScienceLab",
		Author:      "PhysicsTeacher",
	},
	{
		Title:       "Fascinating lab accident today",
		Description: "When two samples got mixed up, we discovered something unexpected...",
		Comments:    178,
		Likes:       456,
		PondName:    "ScienceLab",
		Author:      "AccidentalScientist",
	},
	{
		Title:       "Science communication tips",
		Description: "How to explain complex concepts to non-scientists without dumbing it down",
		Comments:    234,
		Likes:       678,
		PondName:    "ScienceLab",
		Author:      "SciComm",
	},
	{
		Title:       "DIY lab equipment hacks",
		Description: "Budget-friendly solutions that actually work. #4 will surprise you",
		Comments:    145,
		Likes:       567,
		PondName:    "ScienceLab",
		Author:      "FrugalScientist",
	},
	{
		Title:       "The grant writing struggle",
		Description: "Just submitted my first NSF grant. Here's what I learned...",
		Comments:    167,
		Likes:       445,
		PondName:    "ScienceLab",
		Author:      "GrantWriter",
	},
    {
        Title:       "Climate Science Explained",
        Description: "Breaking down complex climate data into understandable insights.",
        Comments:    198,
        Likes:       445,
        PondName:    "ScienceLab",
        Author:      "ClimateScholar",
    },
    {
        Title:       "Neuroscience Breakthroughs",
        Description: "Recent discoveries about how our brains work and what it means.",
        Comments:    167,
        Likes:       389,
        PondName:    "ScienceLab",
        Author:      "BrainExpert",
    },
    {
        Title:       "Genetics 101",
        Description: "Understanding DNA, genes, and recent advances in genetic research.",
        Comments:    178,
        Likes:       412,
        PondName:    "ScienceLab",
        Author:      "GeneGenius",
    },
}

// Update the existing techPosts array
var techPosts = []Post{
    commonTechPost,
    randomTechPost,
    {
        Title:       "Cloud Architecture Patterns",
        Description: "Best practices for designing scalable cloud applications.",
        Comments:    189,
        Likes:       445,
        PondName:    "TechTalk",
        Author:      "CloudArchitect",
    },
	{
		Title:       "Docker is driving me CRAZY - help please!",
		Description: "Ok so I've spent 6 hours trying to get this container working and I'm losing my mind. The ports are mapped but...",
		Comments:    34,
		Likes:       67,
		PondName:    "TechTalk",
		Author:      "DockerPain",
	},
	{
		Title:       "Finally ditched VS Code for Neovim!",
		Description: "Look I know everyone's tired of editor wars but I gotta share this. My productivity has literally doubled since...",
		Comments:    156,
		Likes:       445,
		PondName:    "TechTalk",
		Author:      "VimVeteran",
	},
	{
		Title:       "Anyone else feel like frameworks are getting out of hand?",
		Description: "Just venting here but I swear every time I start a new project there's ANOTHER framework I need to learn...",
		Comments:    287,
		Likes:       892,
		PondName:    "TechTalk",
		Author:      "ByteMaster",
	},
	{
		Title:       "My 2 cents on Rust after 6 months",
		Description: "Coming from Python, here's what surprised me (good and bad). Honestly wish I'd started sooner...",
		Comments:    123,
		Likes:       445,
		PondName:    "TechTalk",
		Author:      "RustNewbie",
	},
	{
		Title:       "PSA: Check your AWS bills folks",
		Description: "Just got hit with a $500 bill because I forgot about a test instance. Here's how to set up billing alerts...",
		Comments:    445,
		Likes:       1023,
		PondName:    "TechTalk",
		Author:      "CloudMaster",
	},
    {
        Title:       "Cybersecurity Essentials",
        Description: "Protecting your applications from common security threats.",
        Comments:    167,
        Likes:       389,
        PondName:    "TechTalk",
        Author:      "SecurityPro",
    },
    {
        Title:       "Machine Learning Projects",
        Description: "Hands-on ML projects for beginners with Python and TensorFlow.",
        Comments:    234,
        Likes:       567,
        PondName:    "TechTalk",
        Author:      "MLEngineer",
    },
    {
        Title:       "Web Performance Tips",
        Description: "Optimizing your web applications for speed and efficiency.",
        Comments:    156,
        Likes:       334,
        PondName:    "TechTalk",
        Author:      "SpeedOptimizer",
    },
}

// Update the existing codingPosts array
var codingPosts = []Post{
    {
        Title:       "Clean Code Principles",
        Description: "Writing maintainable and efficient code that others can understand.",
        Comments:    198,
        Likes:       445,
        PondName:    "CodingPond",
        Author:      "CodeCleaner",
    },
    {
        Title:       "Design Patterns in Go",
        Description: "Implementing common design patterns in Go with practical examples.",
        Comments:    167,
        Likes:       389,
        PondName:    "CodingPond",
        Author:      "GoArchitect",
    },
    {
        Title:       "Test-Driven Development",
        Description: "A practical guide to TDD with real-world examples.",
        Comments:    178,
        Likes:       412,
        PondName:    "CodingPond",
        Author:      "TestMaster",
    },
    {
        Title:       "Microservices Architecture",
        Description: "Building and deploying microservices with modern tools.",
        Comments:    234,
        Likes:       567,
        PondName:    "CodingPond",
        Author:      "MicroservicesPro",
    },
	{
		Title:       "Clean code vs. clever code - a rant",
		Description: "Just spent 4 hours deciphering a 'clever' one-liner. Please, just write readable code...",
		Comments:    234,
		Likes:       876,
		PondName:    "CodingPond",
		Author:      "CleanCoder",
	},
	{
		Title:       "My first open source contribution!",
		Description: "After months of lurking, I finally submitted a PR to a major project. Here's what I learned...",
		Comments:    156,
		Likes:       567,
		PondName:    "CodingPond",
		Author:      "ByteMaster",
	},
	{
		Title:       "Help with recursion in Go",
		Description: "Trying to implement a tree traversal but getting stack overflow. Code snippet below:",
		Comments:    89,
		Likes:       234,
		PondName:    "CodingPond",
		Author:      "GoNewbie",
	},
	{
		Title:       "Why I still use Vim in 2024",
		Description: "A practical guide to modern Vim workflow with LSP, fuzzy finding, and custom macros.",
		Comments:    345,
		Likes:       789,
		PondName:    "CodingPond",
		Author:      "VimWizard",
	},
	{
		Title:       "Junior dev survival guide",
		Description: "Things I wish I knew when I started. #1: It's okay to not know everything...",
		Comments:    445,
		Likes:       999,
		PondName:    "CodingPond",
		Author:      "SeniorDev",
	},
    {
        Title:       "API Design Best Practices",
        Description: "Creating robust and developer-friendly APIs that scale.",
        Comments:    156,
        Likes:       334,
        PondName:    "CodingPond",
        Author:      "APIDesigner",
    },
}

// Update allPosts to include all new posts
var allPosts = []Post{
    // Official posts
    officialWelcomePost,
    officialUpdatePost,
    officialCommunityPost,
    officialGuidelinesPost,
}

func init() {
    // Append all pond-specific posts to allPosts
    allPosts = append(allPosts, artPosts...)
    allPosts = append(allPosts, bookPosts...)
    allPosts = append(allPosts, foodPosts...)
    allPosts = append(allPosts, gardenPosts...)
    allPosts = append(allPosts, sciencePosts...)
    allPosts = append(allPosts, techPosts...)
    allPosts = append(allPosts, codingPosts...)
}

// Add this function to get all posts
func GetAllPosts() []Post {
    return allPosts
} 