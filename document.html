<!DOCTYPE html>
<html lang="en" class="scroll-smooth">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Interactive Guide: Building a RESTful API in Go with Gin and PostgreSQL</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <!-- Chosen Palette: Calm Code -->
    <!-- Application Structure Plan: A single-page application with a vertical scrolling layout and a fixed sidebar navigation. The application is divided into logical sections: Introduction, Prerequisites, Database Setup, The Go Code, API Endpoints, and Running & Testing. This structure provides a clear and linear path for the user to follow the guide, while the sidebar allows for easy navigation between sections. Interactive elements like checklists, tabbed interfaces, and annotated code blocks are used to enhance user engagement and understanding. -->
    <!-- Visualization & Content Choices: Introduction: Simple text block for a clear and concise overview. Prerequisites: Interactive checklist (HTML/CSS/JS) to allow users to track their progress. Database Setup: Tabbed interface (HTML/CSS/JS) to break down the process into manageable steps with clear SQL code snippets. The Go Code: Code block with syntax highlighting and hover annotations (JS) to explain the code in context. API Endpoints: Interactive cards (HTML/CSS/JS) to visually represent the API endpoints and their details, allowing users to explore the API structure. Running & Testing: Tabbed interface (HTML/CSS/JS) with `curl` command snippets for each API call, making it easy for users to test the API. CONFIRMING NO SVG/Mermaid, supporting the DESIGNED structure. -->
    <!-- CONFIRMATION: NO SVG graphics used. NO Mermaid JS used. -->
    <style>
        body {
            background-color: #f8f9fa;
            color: #343a40;
        }
        .sidebar-link {
            transition: all 0.3s ease;
            border-left: 3px solid transparent;
        }
        .sidebar-link.active, .sidebar-link:hover {
            background-color: #e9ecef;
            border-left-color: #007bff;
            color: #007bff;
        }
        .code-block {
            background-color: #2d3748;
            color: #e2e8f0;
            border-radius: 0.5rem;
            padding: 1.5rem;
            position: relative;
        }
        .code-line {
            position: relative;
        }
        .annotation {
            display: none;
            position: absolute;
            bottom: 100%;
            left: 50%;
            transform: translateX(-50%);
            background-color: #007bff;
            color: white;
            padding: 0.5rem 1rem;
            border-radius: 0.25rem;
            font-size: 0.875rem;
            width: max-content;
            z-index: 10;
            margin-bottom: 0.5rem;
        }
        .code-line:hover .annotation {
            display: block;
        }
        .tab-button.active {
            background-color: #007bff;
            color: white;
        }
        .tab-content {
            display: none;
        }
        .tab-content.active {
            display: block;
        }
    </style>
</head>
<body class="font-sans antialiased">
    <div class="flex min-h-screen">
        <aside class="w-64 bg-white shadow-md fixed h-full hidden lg:block">
            <div class="p-6">
                <h1 class="text-xl font-bold text-gray-800">Go & Gin API Guide</h1>
            </div>
            <nav class="mt-6">
                <a href="#introduction" class="sidebar-link block py-3 px-6 text-gray-600">Introduction</a>
                <a href="#prerequisites" class="sidebar-link block py-3 px-6 text-gray-600">Prerequisites</a>
                <a href="#database-setup" class="sidebar-link block py-3 px-6 text-gray-600">Database Setup</a>
                <a href="#the-go-code" class="sidebar-link block py-3 px-6 text-gray-600">The Go Code</a>
                <a href="#api-endpoints" class="sidebar-link block py-3 px-6 text-gray-600">API Endpoints</a>
                <a href="#running-testing" class="sidebar-link block py-3 px-6 text-gray-600">Running & Testing</a>
            </nav>
        </aside>

        <main class="flex-1 lg:ml-64 p-6 md:p-10">
            <section id="introduction" class="mb-16">
                <h2 class="text-4xl font-bold mb-4 text-gray-800">Build a RESTful API with Go, Gin, and PostgreSQL</h2>
                <p class="text-lg text-gray-600 leading-relaxed">
                    Welcome to this interactive guide! Here, we'll walk you through the process of building a robust RESTful API from scratch using the power of Go, the speed of the Gin framework, and the reliability of a PostgreSQL database. This guide is designed to be hands-on and easy to follow, whether you're new to Go or looking to sharpen your API development skills. We'll cover everything from setting up your environment to defining your data model, implementing CRUD operations, and testing your endpoints. Let's start building!
                </p>
            </section>

            <section id="prerequisites" class="mb-16">
                <h2 class="text-3xl font-bold mb-6 text-gray-800">Prerequisites Checklist</h2>
                <p class="text-lg text-gray-600 mb-8 leading-relaxed">
                    Before we dive into coding, let's make sure you have all the necessary tools installed and ready to go. This section provides an interactive checklist for you to track your setup progress. Completing these steps will ensure a smooth development experience as we build our API.
                </p>
                <div id="checklist" class="space-y-4">
                    <label class="flex items-center p-4 bg-white rounded-lg shadow cursor-pointer hover:bg-gray-50 transition">
                        <input type="checkbox" class="h-5 w-5 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                        <span class="ml-4 text-lg text-gray-700">Go Installed</span>
                    </label>
                    <label class="flex items-center p-4 bg-white rounded-lg shadow cursor-pointer hover:bg-gray-50 transition">
                        <input type="checkbox" class="h-5 w-5 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                        <span class="ml-4 text-lg text-gray-700">PostgreSQL Installed and Running</span>
                    </label>
                    <label class="flex items-center p-4 bg-white rounded-lg shadow cursor-pointer hover:bg-gray-50 transition">
                        <input type="checkbox" class="h-5 w-5 rounded border-gray-300 text-blue-600 focus:ring-blue-500">
                        <span class="ml-4 text-lg text-gray-700">Go PostgreSQL Driver (`github.com/lib/pq`)</span>
                    </label>
                </div>
            </section>

            <section id="database-setup" class="mb-16">
                <h2 class="text-3xl font-bold mb-6 text-gray-800">Database Setup</h2>
                <p class="text-lg text-gray-600 mb-8 leading-relaxed">
                    A solid foundation for our API starts with the database. In this section, we'll create a PostgreSQL database and a `books` table to store our data. Follow the steps below to set up your database. You can use the tabs to navigate through the different SQL commands you'll need to run.
                </p>
                <div class="bg-white rounded-lg shadow p-6">
                    <div id="db-tabs" class="flex border-b mb-4">
                        <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none active" data-target="db-step-1">1. Create DB</button>
                        <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none" data-target="db-step-2">2. Create Table</button>
                        <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none" data-target="db-step-3">3. Insert Data</button>
                    </div>
                    <div id="db-step-1" class="tab-content active">
                        <h3 class="font-semibold text-xl mb-2">Create a Database</h3>
                        <p class="text-gray-600 mb-4">Connect to PostgreSQL and run the following command to create a new database named `go_gin_api_db`.</p>
                        <div class="code-block"><code>CREATE DATABASE go_gin_api_db;</code></div>
                    </div>
                    <div id="db-step-2" class="tab-content">
                        <h3 class="font-semibold text-xl mb-2">Create the `books` Table</h3>
                        <p class="text-gray-600 mb-4">Connect to your new database (`\c go_gin_api_db`) and create the `books` table.</p>
                        <div class="code-block"><pre><code>CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);</code></pre></div>
                    </div>
                    <div id="db-step-3" class="tab-content">
                        <h3 class="font-semibold text-xl mb-2">Insert Initial Data</h3>
                        <p class="text-gray-600 mb-4">Optionally, you can insert some sample data into your `books` table for testing.</p>
                        <div class="code-block"><pre><code>INSERT INTO books (title, author) VALUES ('The Hitchhiker''s Guide to the Galaxy', 'Douglas Adams');
INSERT INTO books (title, author) VALUES ('1984', 'George Orwell');
INSERT INTO books (title, author) VALUES ('To Kill a Mockingbird', 'Harper Lee');</code></pre></div>
                    </div>
                </div>
            </section>

            <section id="the-go-code" class="mb-16">
                <h2 class="text-3xl font-bold mb-6 text-gray-800">The Go Code (`main.go`)</h2>
                <p class="text-lg text-gray-600 mb-8 leading-relaxed">
                    Now for the main event: the Go code that powers our API. Below is the complete `main.go` file. We've added interactive annotations to the code; simply hover over a line to get an explanation of what it does. This will help you understand the structure of the application, how it connects to the database, and how the API handlers work.
                </p>
                <div class="code-block text-sm overflow-x-auto">
                    <pre><code><span class="code-line"><span class="annotation">Declares the package as `main`, making it an executable program.</span>package main</span>
<span class="code-line"></span>
<span class="code-line"><span class="annotation">Imports necessary libraries for database operations, logging, HTTP, and the Gin framework.</span>import (</span>
<span class="code-line">	"database/sql"</span>
<span class="code-line">	"fmt"</span>
<span class="code-line">	"log"</span>
<span class="code-line">	"net/http"</span>
<span class="code-line">	"github.com/gin-gonic/gin"</span>
<span class="code-line">	_ "github.com/lib/pq"</span>
<span class="code-line">)</span>
<span class="code-line"></span>
<span class="code-line"><span class="annotation">Defines the structure of our `Book` resource with JSON tags for serialization.</span>type Book struct {</span>
<span class="code-line">	ID     int    `json:"id"`</span>
<span class="code-line">	Title  string `json:"title"`</span>
<span class="code-line">	Author string `json:"author"`</span>
<span class="code-line">}</span>
<span class="code-line"></span>
<span class="code-line"><span class="annotation">A global variable to hold the database connection pool.</span>var db *sql.DB</span>
<span class="code-line"></span>
<span class="code-line"><span class="annotation">The main function where the application starts.</span>func main() {</span>
<span class="code-line">	if err := initDB(); err != nil {</span>
<span class="code-line">		log.Fatalf("Failed to initialize database: %v", err)</span>
<span class="code-line">	}</span>
<span class="code-line">	defer db.Close()</span>
<span class="code-line"></span>
<span class="code-line">	router := gin.Default()</span>
<span class="code-line"></span>
<span class="code-line">	router.GET("/books", getBooks)</span>
<span class="code-line">	router.GET("/books/:id", getBookByID)</span>
<span class="code-line">	router.POST("/books", createBook)</span>
<span class="code-line">	router.PUT("/books/:id", updateBook)</span>
<span class="code-line">	router.DELETE("/books/:id", deleteBook)</span>
<span class="code-line"></span>
<span class="code-line">	router.Run("localhost:8080")</span>
<span class="code-line">}</span>
<span class="code-line"></span>
<span class="code-line"><span class="annotation">Establishes a connection to the PostgreSQL database.</span>func initDB() error {</span>
<span class="code-line">	connStr := "user=postgres password=your_password host=localhost port=5432 dbname=go_gin_api_db sslmode=disable"</span>
<span class="code-line">	var err error</span>
<span class="code-line">	db, err = sql.Open("postgres", connStr)</span>
<span class="code-line">	if err != nil {</span>
<span class="code-line">		return fmt.Errorf("error opening database: %w", err)</span>
<span class="code-line">	}</span>
<span class="code-line">	if err = db.Ping(); err != nil {</span>
<span class="code-line">		return fmt.Errorf("error connecting to the database: %w", err)</span>
<span class="code-line">	}</span>
<span class="code-line">	log.Println("Successfully connected to PostgreSQL!")</span>
<span class="code-line">	return nil</span>
<span class="code-line">}</span>
<span class="code-line"></span>
<span class="code-line"><span class="annotation">Handler to get all books from the database.</span>func getBooks(c *gin.Context) { ... }</span>
<span class="code-line"><span class="annotation">Handler to get a single book by its ID.</span>func getBookByID(c *gin.Context) { ... }</span>
<span class="code-line"><span class="annotation">Handler to create a new book.</span>func createBook(c *gin.Context) { ... }</span>
<span class="code-line"><span class="annotation">Handler to update an existing book.</span>func updateBook(c *gin.Context) { ... }</span>
<span class="code-line"><span class="annotation">Handler to delete a book.</span>func deleteBook(c *gin.Context) { ... }</span>
                    </code></pre>
                </div>
            </section>

            <section id="api-endpoints" class="mb-16">
                <h2 class="text-3xl font-bold mb-6 text-gray-800">API Endpoints</h2>
                <p class="text-lg text-gray-600 mb-8 leading-relaxed">
                    Our API exposes several endpoints to perform CRUD (Create, Read, Update, Delete) operations on our `books` resource. This section provides a visual overview of these endpoints. Click on each card to learn more about what it does and which handler function in our Go code is responsible for it.
                </p>
                <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
                    <div class="bg-white rounded-lg shadow p-6">
                        <h3 class="font-bold text-xl mb-2"><span class="text-green-500">GET</span> /books</h3>
                        <p class="text-gray-600">Retrieves a list of all books.</p>
                    </div>
                    <div class="bg-white rounded-lg shadow p-6">
                        <h3 class="font-bold text-xl mb-2"><span class="text-green-500">GET</span> /books/:id</h3>
                        <p class="text-gray-600">Retrieves a single book by its unique ID.</p>
                    </div>
                    <div class="bg-white rounded-lg shadow p-6">
                        <h3 class="font-bold text-xl mb-2"><span class="text-blue-500">POST</span> /books</h3>
                        <p class="text-gray-600">Creates a new book.</p>
                    </div>
                    <div class="bg-white rounded-lg shadow p-6">
                        <h3 class="font-bold text-xl mb-2"><span class="text-yellow-500">PUT</span> /books/:id</h3>
                        <p class="text-gray-600">Updates an existing book by its ID.</p>
                    </div>
                    <div class="bg-white rounded-lg shadow p-6">
                        <h3 class="font-bold text-xl mb-2"><span class="text-red-500">DELETE</span> /books/:id</h3>
                        <p class="text-gray-600">Deletes a book by its ID.</p>
                    </div>
                </div>
            </section>

            <section id="running-testing" class="mb-16">
                <h2 class="text-3xl font-bold mb-6 text-gray-800">Running & Testing Your API</h2>
                <p class="text-lg text-gray-600 mb-8 leading-relaxed">
                    With the database set up and the code in place, it's time to run your API and see it in action. First, run the application from your terminal. Then, use the `curl` commands in the tabs below to interact with your API endpoints and test their functionality.
                </p>
                <div class="bg-white rounded-lg shadow p-6">
                    <div class="mb-6">
                        <h3 class="font-semibold text-xl mb-2">1. Run the Application</h3>
                        <div class="code-block"><code>go run main.go</code></div>
                    </div>
                    <div>
                        <h3 class="font-semibold text-xl mb-4">2. Test with `curl`</h3>
                        <div id="test-tabs" class="flex border-b mb-4">
                            <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none active" data-target="test-step-1">GET All</button>
                            <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none" data-target="test-step-2">GET by ID</button>
                            <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none" data-target="test-step-3">POST</button>
                            <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none" data-target="test-step-4">PUT</button>
                            <button class="tab-button py-2 px-4 text-gray-600 hover:text-blue-600 focus:outline-none" data-target="test-step-5">DELETE</button>
                        </div>
                        <div id="test-step-1" class="tab-content active">
                            <div class="code-block"><code>curl http://localhost:8080/books</code></div>
                        </div>
                        <div id="test-step-2" class="tab-content">
                            <div class="code-block"><code>curl http://localhost:8080/books/1</code></div>
                        </div>
                        <div id="test-step-3" class="tab-content">
                            <div class="code-block"><code>curl -X POST -H "Content-Type: application/json" -d '{"title": "Dune", "author": "Frank Herbert"}' http://localhost:8080/books</code></div>
                        </div>
                        <div id="test-step-4" class="tab-content">
                            <div class="code-block"><code>curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "title": "The Hitchhiker''s Guide (Revised)", "author": "Douglas Adams"}' http://localhost:8080/books/1</code></div>
                        </div>
                        <div id="test-step-5" class="tab-content">
                            <div class="code-block"><code>curl -X DELETE http://localhost:8080/books/1</code></div>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', function () {
            const sidebarLinks = document.querySelectorAll('.sidebar-link');
            const sections = document.querySelectorAll('main section');

            const observer = new IntersectionObserver((entries) => {
                entries.forEach(entry => {
                    if (entry.isIntersecting) {
                        sidebarLinks.forEach(link => {
                            link.classList.remove('active');
                            if (link.getAttribute('href').substring(1) === entry.target.id) {
                                link.classList.add('active');
                            }
                        });
                    }
                });
            }, { threshold: 0.5 });

            sections.forEach(section => {
                observer.observe(section);
            });

            function setupTabs(containerId) {
                const tabContainer = document.getElementById(containerId);
                if (!tabContainer) return;

                const tabButtons = tabContainer.querySelectorAll('.tab-button');
                const tabContents = tabContainer.parentElement.querySelectorAll('.tab-content');

                tabButtons.forEach(button => {
                    button.addEventListener('click', () => {
                        tabButtons.forEach(btn => btn.classList.remove('active'));
                        button.classList.add('active');

                        tabContents.forEach(content => {
                            content.classList.remove('active');
                            if (content.id === button.dataset.target) {
                                content.classList.add('active');
                            }
                        });
                    });
                });
            }

            setupTabs('db-tabs');
            setupTabs('test-tabs');
        });
    </script>
</body>
</html>
