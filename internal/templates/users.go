package templates

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "net/http"
    "sort"
    "strconv"
    "time"
)

var (
    commonTechPost = Post{
        Title:       "The Future of AI Development",
        Description: "Discussing the latest breakthroughs in artificial intelligence and what they mean for developers.",
        Comments:    156,
        Likes:       432,
        PondName:    "TechTalk",
        Author:      "ByteMaster",
        SecondsAgo:  rand.Intn(2592000), // Random time up to 30 days ago
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
        Description: "We're excited to have you join our community...",
        Comments:    789,
        Likes:       1234,
        PondName:    "Official",
        Author:      "Ribbit Admin",
        SecondsAgo:  rand.Intn(31536000),
    }

    officialUpdatePost = Post{
        Title:       "New Features Coming Soon",
        Description: "We're working on some exciting updates...",
        Comments:    456,
        Likes:       987,
        PondName:    "Official",
        Author:      "Ribbit Admin",
        SecondsAgo:  rand.Intn(31536000),
    }

    musicPond = Pond{Name: "MusicLounge", Description: "For music lovers and creators", Members: "12.3K"}
    gamingPond = Pond{Name: "GamerHaven", Description: "Gaming discussions and news", Members: "18.7K"}
    filmPond = Pond{Name: "CinemaSpot", Description: "Movie reviews and discussions", Members: "14.2K"}
    petsPond = Pond{Name: "PetPals", Description: "All about our furry friends", Members: "10.1K"}
    fitnessPond = Pond{Name: "FitnessZone", Description: "Health and workout tips", Members: "13.4K"}
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
    Author      string
    SecondsAgo  int
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
    var posts []Post
    
    // Get all available posts for user's ponds
    for _, pond := range userPonds {
        switch pond.Name {
        case "TechTalk":
            posts = append(posts, techPosts...)
        case "GreenThumb":
            posts = append(posts, gardenPosts...)
        case "ArtistsCorner":
            posts = append(posts, artPosts...)
        case "BookClub":
            posts = append(posts, bookPosts...)
        case "FoodiesUnite":
            posts = append(posts, foodPosts...)
        case "ScienceLab":
            posts = append(posts, sciencePosts...)
        case "CodingPond":
            posts = append(posts, codingPosts...)
        case "MusicLounge":
            posts = append(posts, musicPosts...)
        case "GamerHaven":
            posts = append(posts, gamingPosts...)
        case "CinemaSpot":
            posts = append(posts, filmPosts...)
        case "PetPals":
            posts = append(posts, petsPosts...)
        case "FitnessZone":
            posts = append(posts, fitnessPosts...)
        }
    }

    // Filter out posts by the current user
    filteredPosts := make([]Post, 0)
    for _, post := range posts {
        if post.Author != excludeAuthor {
            filteredPosts = append(filteredPosts, post)
        }
    }

    // Shuffle the posts
    rand.Shuffle(len(filteredPosts), func(i, j int) {
        filteredPosts[i], filteredPosts[j] = filteredPosts[j], filteredPosts[i]
    })

    return filteredPosts
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
        {Name: "MusicLounge", Description: "For music lovers and creators", Members: "12.3K"},
        {Name: "GamerHaven", Description: "Gaming discussions and news", Members: "18.7K"},
        {Name: "CinemaSpot", Description: "Movie reviews and discussions", Members: "14.2K"},
        {Name: "PetPals", Description: "All about our furry friends", Members: "10.1K"},
        {Name: "FitnessZone", Description: "Health and workout tips", Members: "13.4K"},
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
        {Name: "CodingPond", Description: "Programming discussions", Members: "12.8K"},
        {Name: "ScienceLab", Description: "Scientific discoveries", Members: "11.2K"},
        {Name: "GamerHaven", Description: "Gaming discussions and news", Members: "18.7K"},
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
        {Name: "CinemaSpot", Description: "Movie reviews and discussions", Members: "14.2K"},
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

// GetTrendingPosts returns the top 8 posts by engagement (likes + comments)
func GetTrendingPosts() []Post {
    // Create a copy of allPosts to sort
    posts := make([]Post, len(allPosts))
    copy(posts, allPosts)

    // Sort posts by total engagement (likes + comments)
    sort.Slice(posts, func(i, j int) bool {
        engagementI := posts[i].Likes + posts[i].Comments
        engagementJ := posts[j].Likes + posts[j].Comments
        return engagementI > engagementJ
    })

    // Return only top 8 posts
    if len(posts) > 8 {
        return posts[:8]
    }
    return posts
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
    {
        Title:       "AI Art: Friend or Foe?",
        Description: "A balanced discussion on how AI is impacting the art community.",
        Comments:    789,
        Likes:       1876,
        PondName:    "ArtistsCorner",
        Author:      "ArtPhilosopher",
    },
    {
        Title:       "Urban Sketching Guide",
        Description: "Tips and techniques for capturing city life in your sketchbook.",
        Comments:    345,
        Likes:       987,
        PondName:    "ArtistsCorner",
        Author:      "UrbanArtist",
    },
    {
        Title:       "Color Theory Deep Dive",
        Description: "Understanding color relationships for more impactful artwork.",
        Comments:    234,
        Likes:       876,
        PondName:    "ArtistsCorner",
        Author:      "ColorMaster",
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
    {
        Title:       "The Return of Short Stories",
        Description: "Why short story collections are making a comeback in modern literature.",
        Comments:    432,
        Likes:       1123,
        PondName:    "BookClub",
        Author:      "StoryTeller",
    },
    {
        Title:       "Fantasy vs Science Fiction",
        Description: "The blurring lines between fantasy and sci-fi in contemporary literature.",
        Comments:    345,
        Likes:       876,
        PondName:    "BookClub",
        Author:      "GenreJunkie",
    },
    {
        Title:       "Independent Bookstores Revival",
        Description: "How local bookshops are thriving in the digital age.",
        Comments:    234,
        Likes:       765,
        PondName:    "BookClub",
        Author:      "BookKeeper",
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
    {
        Title:       "Fermentation Revolution",
        Description: "Getting started with home fermentation: kombucha, kimchi, and beyond.",
        Comments:    567,
        Likes:       1432,
        PondName:    "FoodiesUnite",
        Author:      "FermentationFanatic",
    },
    {
        Title:       "Zero-Waste Cooking Guide",
        Description: "Creative ways to use every part of your ingredients.",
        Comments:    432,
        Likes:       1234,
        PondName:    "FoodiesUnite",
        Author:      "SustainableChef",
    },
    {
        Title:       "Regional Pasta Shapes",
        Description: "The history and purpose behind different pasta shapes across Italy.",
        Comments:    345,
        Likes:       987,
        PondName:    "FoodiesUnite",
        Author:      "PastaMaster",
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
    {
        Title:       "Vertical Gardening Solutions",
        Description: "Making the most of limited space with vertical gardens.",
        Comments:    432,
        Likes:       1234,
        PondName:    "GreenThumb",
        Author:      "VerticalGrower",
    },
    {
        Title:       "Native Plant Revolution",
        Description: "Why and how to incorporate native species into your garden.",
        Comments:    345,
        Likes:       987,
        PondName:    "GreenThumb",
        Author:      "NativePlanter",
    },
    {
        Title:       "Companion Planting Guide",
        Description: "Boost your garden's health with strategic plant partnerships.",
        Comments:    234,
        Likes:       876,
        PondName:    "GreenThumb",
        Author:      "GardenSynergist",
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
    {
        Title:       "Black Holes: New Discoveries",
        Description: "Recent observations have challenged our understanding of black hole behavior.",
        Comments:    876,
        Likes:       2341,
        PondName:    "ScienceLab",
        Author:      "CosmicMind",
    },
    {
        Title:       "CRISPR Breakthroughs 2024",
        Description: "The latest developments in gene editing technology and their implications.",
        Comments:    654,
        Likes:       1567,
        PondName:    "ScienceLab",
        Author:      "GeneGenius",
    },
    {
        Title:       "Ocean Floor Mysteries",
        Description: "New species discovered in deep-sea exploration mission. Pictures inside!",
        Comments:    432,
        Likes:       1234,
        PondName:    "ScienceLab",
        Author:      "DeepDiver",
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
    {
        Title:       "The Rise of Quantum Computing",
        Description: "Breaking down the latest advancements in quantum computing and what it means for developers.",
        Comments:    567,
        Likes:       1432,
        PondName:    "TechTalk",
        Author:      "QuantumCoder",
    },
    {
        Title:       "Web Assembly: The Future?",
        Description: "How WASM is changing the landscape of web development. Real-world examples inside.",
        Comments:    432,
        Likes:       987,
        PondName:    "TechTalk",
        Author:      "WebWizard",
    },
    {
        Title:       "My Journey into Embedded Systems",
        Description: "From web dev to embedded systems - lessons learned and pitfalls to avoid.",
        Comments:    234,
        Likes:       654,
        PondName:    "TechTalk",
        Author:      "ChipMaster",
    },
}

// Update the existing codingPosts array
var codingPosts = []Post{
    {
        Title:       "The Art of Clean Code",
        Description: "Best practices for writing maintainable and efficient code.",
        Comments:    234,
        Likes:       567,
        PondName:    "CodingPond",
        Author:      "ByteMaster",
        SecondsAgo:  rand.Intn(31536000), // Random time up to 1 year ago (365 days)
    },
    {
        Title:       "Modern Vim Workflow",
        Description: "A practical guide to modern Vim workflow with LSP, fuzzy finding, and custom macros.",
        Comments:    345,
        Likes:       789,
        PondName:    "CodingPond",
        Author:      "VimWizard",
        SecondsAgo:  rand.Intn(31536000),
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
    {
        Title:       "Rust vs Go in 2024",
        Description: "Comparing two modern languages for system programming.",
        Comments:    789,
        Likes:       1987,
        PondName:    "CodingPond",
        Author:      "SystemsGuru",
    },
    {
        Title:       "Microservices: The Good Parts",
        Description: "Real-world lessons from 5 years of microservice architecture.",
        Comments:    654,
        Likes:       1654,
        PondName:    "CodingPond",
        Author:      "MicroMaster",
    },
    {
        Title:       "Frontend Testing Evolution",
        Description: "How frontend testing has evolved with modern frameworks.",
        Comments:    432,
        Likes:       1234,
        PondName:    "CodingPond",
        Author:      "TestingPro",
    },
}

// Update new post arrays for each new pond
var musicPosts = []Post{
    {
        Title:       "Learning Jazz Piano - Month 1",
        Description: "Started my jazz piano journey. Here's what I've learned about chord progressions...",
        Comments:    145,
        Likes:       432,
        PondName:    "MusicLounge",
        Author:      "JazzMaster",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Best DAWs for Beginners 2024",
        Description: "Comparing popular Digital Audio Workstations for newcomers to music production.",
        Comments:    234,
        Likes:       567,
        PondName:    "MusicLounge",
        Author:      "BeatMaker",
    },
    {
        Title:       "Classical vs Jazz: The Theory Overlap",
        Description: "Interesting connections between classical and jazz music theory that might surprise you.",
        Comments:    189,
        Likes:       445,
        PondName:    "MusicLounge",
        Author:      "TheoryNerd",
    },
    {
        Title:       "Why I switched back to vinyl",
        Description: "Digital is convenient, but there's something about analog that just hits different...",
        Comments:    234,
        Likes:       678,
        PondName:    "MusicLounge",
        Author:      "VinylVeteran",
    },
    {
        Title:       "Music theory is clicking finally!",
        Description: "After 6 months of study, intervals and chord progressions are starting to make sense...",
        Comments:    123,
        Likes:       345,
        PondName:    "MusicLounge",
        Author:      "MusicNewbie",
    },
    {
        Title:       "Studio Monitor Recommendations?",
        Description: "Budget is around $500, mainly for electronic music production. Currently looking at...",
        Comments:    167,
        Likes:       234,
        PondName:    "MusicLounge",
        Author:      "BeatMaker",
    },
}

var gamingPosts = []Post{
    {
        Title:       "The Rise of Indie Games",
        Description: "How independent developers are reshaping the gaming industry...",
        Comments:    345,
        Likes:       789,
        PondName:    "GamerHaven",
        Author:      "IndieGamer",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Next-Gen Console Comparison",
        Description: "Detailed analysis of the latest gaming consoles: specs, games, and value.",
        Comments:    567,
        Likes:       1023,
        PondName:    "GamerHaven",
        Author:      "TechGamer",
    },
    {
        Title:       "The Psychology of Game Design",
        Description: "How game developers use psychology to create engaging experiences.",
        Comments:    234,
        Likes:       678,
        PondName:    "GamerHaven",
        Author:      "GamePsych",
    },
    {
        Title:       "Finally beat Malenia!",
        Description: "After 47 attempts, I finally did it! Here's the build that worked for me...",
        Comments:    445,
        Likes:       1289,
        PondName:    "GamerHaven",
        Author:      "EldenLord",
    },
    {
        Title:       "The state of game preservation",
        Description: "With digital-only releases and server shutdowns, we're losing gaming history...",
        Comments:    234,
        Likes:       876,
        PondName:    "GamerHaven",
        Author:      "RetroGamer",
    },
    {
        Title:       "My first speedrun experience",
        Description: "Decided to try speedrunning Hollow Knight. The community is amazing...",
        Comments:    156,
        Likes:       567,
        PondName:    "GamerHaven",
        Author:      "SpeedRunner",
    },
}

var filmPosts = []Post{
    {
        Title:       "The Evolution of CGI in Cinema",
        Description: "From Jurassic Park to today: How CGI has transformed moviemaking.",
        Comments:    278,
        Likes:       567,
        PondName:    "CinemaSpot",
        Author:      "FilmBuff",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Hidden Gems: Underrated Films of 2023",
        Description: "Great movies you might have missed this year.",
        Comments:    189,
        Likes:       445,
        PondName:    "CinemaSpot",
        Author:      "CinematicArt",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "The lost art of practical effects",
        Description: "Modern CGI is amazing, but there's something special about practical effects...",
        Comments:    345,
        Likes:       789,
        PondName:    "CinemaSpot",
        Author:      "FilmNoir",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Best Director's Cuts?",
        Description: "Sometimes longer is better. Here are some films that were improved by their director's cut...",
        Comments:    234,
        Likes:       567,
        PondName:    "CinemaSpot",
        Author:      "CinematicArt",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Sound design appreciation thread",
        Description: "Let's talk about films with outstanding sound design. Starting with Dune...",
        Comments:    189,
        Likes:       456,
        PondName:    "CinemaSpot",
        Author:      "SoundGeek",
        SecondsAgo:  rand.Intn(31536000),
    },
}

var petsPosts = []Post{
    {
        Title:       "Understanding Cat Body Language",
        Description: "A comprehensive guide to what your cat is trying to tell you.",
        Comments:    234,
        Likes:       678,
        PondName:    "PetPals",
        Author:      "CatWhisperer",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "First-Time Dog Owner Guide",
        Description: "Everything you need to know before getting your first dog.",
        Comments:    345,
        Likes:       789,
        PondName:    "PetPals",
        Author:      "DogTrainer",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "My rescue story",
        Description: "After months of patience, my anxious rescue dog finally trusts me. Here's what worked...",
        Comments:    567,
        Likes:       1432,
        PondName:    "PetPals",
        Author:      "RescueHero",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Raw diet myths debunked",
        Description: "Veterinarian here! Let's clear up some misconceptions about raw feeding...",
        Comments:    345,
        Likes:       876,
        PondName:    "PetPals",
        Author:      "VetDoc",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Apartment-friendly pets",
        Description: "Not just cats and dogs! Here are some great pets for small living spaces...",
        Comments:    234,
        Likes:       567,
        PondName:    "PetPals",
        Author:      "UrbanPets",
        SecondsAgo:  rand.Intn(31536000),
    },
}

var fitnessPosts = []Post{
    {
        Title:       "Myth-Busting: Common Fitness Misconceptions",
        Description: "Separating fact from fiction in the fitness world.",
        Comments:    289,
        Likes:       567,
        PondName:    "FitnessZone",
        Author:      "FitCoach",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Building a Home Gym on a Budget",
        Description: "Smart ways to create your workout space without breaking the bank.",
        Comments:    178,
        Likes:       445,
        PondName:    "FitnessZone",
        Author:      "HomeGymPro",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "From couch to 5K - Success!",
        Description: "Just finished my first 5K after being sedentary for years. Here's my journey...",
        Comments:    345,
        Likes:       987,
        PondName:    "FitnessZone",
        Author:      "RunnerNewbie",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "The protein myth",
        Description: "Sports nutritionist here! Let's talk about how much protein you really need...",
        Comments:    456,
        Likes:       1023,
        PondName:    "FitnessZone",
        Author:      "NutritionPro",
        SecondsAgo:  rand.Intn(31536000),
    },
    {
        Title:       "Mobility work changed my life",
        Description: "After years of lifting, adding mobility work made the biggest difference...",
        Comments:    234,
        Likes:       678,
        PondName:    "FitnessZone",
        Author:      "FlexMaster",
        SecondsAgo:  rand.Intn(31536000),
    },
}

// Update official posts array
var officialCommunityPost = Post{
    Title:       "Community Update: New Features",
    Description: "Check out our latest platform improvements and upcoming changes...",
    Comments:    567,
    Likes:       890,
    PondName:    "Official",
    Author:      "Ribbit Admin",
    SecondsAgo:  rand.Intn(31536000),
}

var officialGuidelinesPost = Post{
    Title:       "Community Guidelines Reminder",
    Description: "A friendly reminder about our community standards and how to keep Ribbit welcoming for everyone...",
    Comments:    234,
    Likes:       567,
    PondName:    "Official",
    Author:      "Ribbit Admin",
    SecondsAgo:  rand.Intn(31536000),
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
    rand.Seed(time.Now().UnixNano())
    // Append all pond-specific posts to allPosts
    allPosts = append(allPosts, artPosts...)
    allPosts = append(allPosts, bookPosts...)
    allPosts = append(allPosts, foodPosts...)
    allPosts = append(allPosts, gardenPosts...)
    allPosts = append(allPosts, sciencePosts...)
    allPosts = append(allPosts, techPosts...)
    allPosts = append(allPosts, codingPosts...)
    allPosts = append(allPosts, musicPosts...)
    allPosts = append(allPosts, gamingPosts...)
    allPosts = append(allPosts, filmPosts...)
    allPosts = append(allPosts, petsPosts...)
    allPosts = append(allPosts, fitnessPosts...)
}

// Add this helper function
func formatTimeAgo(secondsAgo int) string {
    minutes := secondsAgo / 60
    hours := minutes / 60
    days := hours / 24
    weeks := days / 7
    months := days / 30

    if months > 0 {
        if months == 1 {
            return "1 month ago"
        }
        return fmt.Sprintf("%d months ago", months)
    }
    if weeks > 0 {
        if weeks == 1 {
            return "1 week ago"
        }
        return fmt.Sprintf("%d weeks ago", weeks)
    }
    if days > 0 {
        if days == 1 {
            return "1 day ago"
        }
        return fmt.Sprintf("%d days ago", days)
    }
    if hours > 0 {
        if hours == 1 {
            return "1 hour ago"
        }
        return fmt.Sprintf("%d hours ago", hours)
    }
    if minutes > 0 {
        if minutes == 1 {
            return "1 minute ago"
        }
        return fmt.Sprintf("%d minutes ago", minutes)
    }
    return "just now"
}

// Modify the GetAllPosts function to sort by time
func GetAllPosts() []Post {
    // Create a copy of allPosts to sort
    posts := make([]Post, len(allPosts))
    copy(posts, allPosts)
    
    // Sort posts by SecondsAgo (most recent first)
    sort.Slice(posts, func(i, j int) bool {
        return posts[i].SecondsAgo < posts[j].SecondsAgo
    })
    
    return posts
}

// Add these new functions
func (u *UserTemplate) GetPaginatedPosts(start, count int) []Post {
    allUserPosts := getRandomPostsForUser(u.Ponds, u.Email)
    
    // Ensure we don't exceed slice bounds
    end := start + count
    if end > len(allUserPosts) {
        end = len(allUserPosts)
    }
    if start >= len(allUserPosts) {
        return []Post{}
    }
    
    return allUserPosts[start:end]
}

// Add this handler function
func HandleGetPosts(w http.ResponseWriter, r *http.Request) {
    fmt.Println("HandleGetPosts called")
    
    // Parse query parameters
    startStr := r.URL.Query().Get("start")
    countStr := r.URL.Query().Get("count")
    
    start, err := strconv.Atoi(startStr)
    if err != nil {
        fmt.Println("Error parsing start:", err)
        start = 0
    }
    
    count, err := strconv.Atoi(countStr)
    if err != nil {
        fmt.Println("Error parsing count:", err)
        count = 3
    }
    
    fmt.Printf("Fetching posts from %d to %d\n", start, count)
    
    // Get user from session
    user, ok := r.Context().Value("user").(*UserTemplate)
    if !ok {
        fmt.Println("User not found in context")
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
    
    // Get paginated posts
    posts := user.GetPaginatedPosts(start, count)
    fmt.Printf("Found %d posts\n", len(posts))
    
    // Set JSON content type
    w.Header().Set("Content-Type", "application/json")
    
    // Encode posts as JSON and send response
    if err := json.NewEncoder(w).Encode(posts); err != nil {
        fmt.Println("Error encoding posts:", err)
        http.Error(w, "Failed to encode posts", http.StatusInternalServerError)
        return
    }
}
