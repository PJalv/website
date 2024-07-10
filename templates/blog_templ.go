// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "strings"

func Header(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><!-- Open Graph Protocol tags --><meta property=\"og:title\" content=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 10, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><meta property=\"og:description\" content=\"Click to view more\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 12, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><meta property=\"og:image\" content=\"https://media.licdn.com/dms/image/D4E03AQFaHGzaKImG2Q/profile-displayphoto-shrink_200_200/0/1663912839285?e=2147483647&amp;v=beta&amp;t=8UaJlQkIIsikNMqBAr_ie_wEMlCzNI3ogrzpoOdXJek\"><link rel=\"icon\" type=\"image/x-icon\" href=\"/favicon.ico\"><script src=\"https://unpkg.com/htmx.org@1.8.5/dist/htmx.min.js\"></script><script src=\"https://unpkg.com/htmx.org/dist/ext/json-enc.js\"></script><style>\n    body {\n      font-family: Arial, sans-serif;\n      margin: 0;\n      padding: 0;\n      background-color: #1a1b26;\n      /* Dark background */\n      color: #c0caf5;\n      /* Light text */\n    }\n\n    .container {\n      max-width: 800px;\n      margin: 30px auto;\n      padding: 20px;\n      background: #24283b;\n      /* Slightly lighter background for container */\n      box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);\n      border-radius: 8px;\n    }\n\n    .container li {\n      padding-top: 5px;\n      padding-bottom: 5px;\n    }\n    .landing li{\n      list-style-type: none;\n    }\n    .navbar {\n      display: flex;\n      justify-content: space-between;\n      align-items: center;\n      max-width: 800px;\n      margin: 20px auto;\n      padding-left: 20px;\n      padding-right: 20px;\n      background: #24283b;\n      /* Slightly lighter background for navbar */\n      box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);\n      border-radius: 8px;\n      flex-direction: row;\n      /* Arrange items from left to right */\n    }\n\n\n\n    #blog-container ul{\n    list-style-type: none;\n    padding-left: 0px;\n    }\n    nav ul {\n      list-style-type: none;\n      padding: 0;\n      display: flex;\n      /* Make the list items flex containers */\n    }\n\n    nav li {\n      margin-right: 20px;\n      /* Adjust this value to control the spacing between links */\n    }\n\n    nav a {\n      color: #bb9af7\n    }\n    .post-content ul{\n      list-style-type: disc;\n      }\n\n\n    nav li:last-child {\n      margin-right: 0;\n      /* Remove margin from the last item to avoid extra spacing */\n    }\n\n    p.blog-date {\n      color: #ff9e64;\n    }\n\n    h1 {\n      text-align: center;\n      color: #7aa2f7;\n      /* Blue header */\n    }\n\n    h2 {\n      color: #bb9af7;\n      /* Purple sub-headers */\n    }\n\n    h3 {\n      color: #7dcfff;\n    }\n\n    h4 {\n      color: #4fd6be;\n    }\n\n    a {\n      color: #7aa2f7;\n      /* Blue links */\n      text-decoration: none;\n    }\n\n    a:hover {\n      text-decoration: underline;\n    }\n\n\n    .intro {\n      display: flex;\n      align-items: center;\n      margin: 20px auto;\n    }\n\n    .profile-images {\n      display: flex;\n      justify-content: center;\n    }\n\n    .profile-images img {\n      border-radius: 50%;\n      width: 130px;\n      height: 130px;\n      margin-right: 20px;\n      box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);\n    }\n\n    .intro-text {\n      font-size: 1.1em;\n      line-height: 1.5em;\n    }\n\n    pre {\n      background-color: #292e42;\n      /* Darker background for code blocks */\n      padding: 15px;\n      border-radius: 8px;\n      overflow: auto;\n      position: relative;\n    }\n\n    code {\n      font-family: 'Courier New', Courier, monospace;\n      color: #9574e1;\n      /* Light text for code */\n      background-color: #3d3d3d;\n      border-radius: 2px;\n      padding-left: 2px;\n      padding-right: 2px;\n      padding-top: 4px;\n    }\n\n    .modal {\n      display: none;\n      position: fixed;\n      z-index: 1;\n      left: 0;\n      top: 0;\n      width: 100%;\n      height: 100%;\n      overflow: auto;\n      background-color: rgba(0, 0, 0, 0.9);\n    }\n\n    .modal-content {\n      margin: 5% auto;\n      display: block;\n      width: 100%;\n      max-width: 900px;\n    }\n\n    .close {\n      position: absolute;\n      top: 10px;\n      right: 25px;\n      color: #fff;\n      font-size: 35px;\n      font-weight: bold;\n      cursor: pointer;\n    }\n\n\n    .modal img {\n      display: flex;\n      justify-content: center;\n      margin-top: 10%;\n      width: 100%;\n    }\n\n    .date-button {\n      display: flex;\n      justify-content: flex-end;\n      margin-top: -1.5em;\n      margin-bottom: 14px;\n    }\n\n    button {\n      background-color: #bb9af7;\n      border: none;\n      border-radius: 5px;\n      box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);\n    }\n\n    video {\n      max-width: 95%;\n      display: block;\n      margin: auto;\n    }\n\n    .copy-button {\n      position: absolute;\n      top: 10px;\n      right: 10px;\n      background-color: #4CAF50;\n      color: white;\n      border: none;\n      padding: 5px 10px;\n      cursor: pointer;\n      border-radius: 5px;\n      font-size: 0.9em;\n    }\n\n    .copy-button:hover {\n      background-color: #45a049;\n    }\n\n    pre::-webkit-scrollbar {\n      width: 8px;\n      height: 8px;\n    }\n\n    pre::-webkit-scrollbar-thumb {\n      background: #3b4261;\n      border-radius: 8px;\n    }\n\n    pre::-webkit-scrollbar-thumb:hover {\n      background: #5e6687;\n    }\n\n    pre::-webkit-scrollbar-track {\n      background: #24283b;\n      border-radius: 8px;\n    }\n\n    .blog-post {\n      max-width: 100%;\n      /* Ensure blog posts don't exceed the width of the container */\n    }\n\n    .blog-post img {\n      max-height: 30em;\n    }\n\n    .blog-post figcaption {\n      font-size: smaller;\n      color: #bb9af7;\n      display: flex;\n      justify-content: center;\n    }\n\n    .postImg {\n      color: #7aa2f7;\n      text-decoration: none;\n      display: flex;\n      justify-content: center;\n    }\n\n    .post-date {\n      margin-left: auto;\n      color: #bb9af7;\n      text-decoration: none;\n    }\n\n    .post-date:hover {\n      text-decoration: none;\n    }\n\n    .description:hover {\n      text-decoration: none;\n    }\n\n    .post-meta {\n      display: flex;\n      align-items: center;\n    }\n\n    .description {\n      margin-top: 5px;\n      font-size: 0.85em;\n      font-style: italic;\n      color: white;\n      text-decoration: none;\n      /* Adjust as needed for spacing */\n    }\n\n    .scrolling-text-container {\n      max-width: 400px;\n      /* adjust as needed */\n      overflow: hidden;\n      white-space: nowrap;\n    }\n\n    .scrolling-text {\n      display: inline-block;\n      animation: scrollText 15s linear infinite;\n    }\n\n    @keyframes scrollText {\n      from {\n        transform: translateX(100%);\n      }\n\n      to {\n        transform: translateX(-100%);\n      }\n    }\n\n    /* Stop animation when text fits container */\n    .scrolling-text-container:hover .scrolling-text {\n      animation-play-state: paused;\n    }\n  </style></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func NavBar() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"navbar\" align=\"center\"><ul><li><a href=\"/\" class=\"navbar-brand\">Home</a></li><li><a href=\"/blog\" class=\"navbar-brand\">Blog</a></li></ul></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Landing() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body><div align=\"center\" class=\"container\"><div class=\"landing\"><h1>Jorge Luis Suarez</h1><div class=\"profile-images\"><img src=\"https://media.licdn.com/dms/image/D4E03AQFaHGzaKImG2Q/profile-displayphoto-shrink_200_200/0/1663912839285?e=2147483647&amp;v=beta&amp;t=8UaJlQkIIsikNMqBAr_ie_wEMlCzNI3ogrzpoOdXJek\" alt=\"Jorge Luis Suarez\"> <img src=\"https://avatars.githubusercontent.com/u/134026493?v=4\" alt=\"PJalv\"></div><div class=\"intro\"><div align=\"left\" class=\"intro-text\"><p>I am a recent graduate with a degree in computer engineering from Cal Poly Pomona, passionate about learning and using new technology to create cool stuff. I am currently really into working on both embedded systems as well as web and cloud development. Here you can find some of my school projects, including my Senior Design Project, among other things I set myself to create.</p></div></div><h2>Personal Links</h2><ul><li><a href=\"/jorge\" target=\"_blank\">LinkedIn</a></li><li><a href=\"https://github.com/pjalv\" target=\"_blank\">GitHub</a></li><li><a href=\"https://x.com/PJalv\" target=\"_blank\">Twitter</a></li></ul><h2>Senior Project Links</h2><ul><li><a href=\"/iot\" target=\"_blank\">Senior Design Project - IoT Device Control System</a></li><li><a href=\"/iot-presentation\" target=\"_blank\">Video Presentation - Symposium Day</a></li></ul><h2>Other Links</h2><ul><li><a href=\"https://www.example.com/other-projects\" target=\"_blank\">Other Projects</a></li></ul></div></div></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CompBlogData(posts []Post, reverse bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"blog-data\"><div class=\"date-button\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if reverse == true {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"button\" hx-trigger=\"click\" hx-post=\"/blog-data-rev\" hx-headers=\"{&#34;reverse&#34;: &#34;false&#34;}\" hx-target=\"#blog-data\" class=\"toggle-date\">Date &uarr;</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"button\" hx-trigger=\"click\" hx-post=\"/blog-data-rev\" hx-headers=\"{&#34;reverse&#34;: &#34;true&#34;}\" hx-target=\"#blog-data\" class=\"toggle-date\">Date &darr;</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, post := range posts {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><div class=\"blog-post-index\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 templ.SafeURL = templ.URL("/blog/" + strings.ToLower(post.Title))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var7)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"post-meta\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strings.ReplaceAll(post.Title, "-", " "))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 430, Col: 49}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <span class=\"post-date\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(post.Date)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 431, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div></a><div class=\"description\">&ensp;")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(post.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 434, Col: 54}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></li>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func BlogIndex(posts []Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var11 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var11 == nil {
			templ_7745c5c3_Var11 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavBar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1>Blog</h1><div class=\"container\" id=\"blog-container\"><ul>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CompBlogData(posts, false).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul></div></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func BlogPage(post Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavBar().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container\"><p class=\"blog-date\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(post.Date)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 457, Col: 35}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><article class=\"blog-post\"><header class=\"post-header\"></header><div class=\"post-content\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.Raw(post.Content).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></article></div><div id=\"myModal\" class=\"modal\" onclick=\"closeModalOnClick(event)\"><span class=\"close\" onclick=\"closeModal()\">&times;</span> <img id=\"modal-img\" class=\"modal-content\" src=\"\" alt=\"Full-size image\"></div><script>\n    function openModal(event, src) {\n      event.preventDefault();\n      var modal = document.getElementById(\"myModal\");\n      var modalImg = document.getElementById(\"modal-img\");\n      modal.style.display = \"block\";\n      modalImg.src = src;\n      document.addEventListener('keydown', closeModalOnEscape);\n    }\n\n    function closeModal() {\n      document.getElementById(\"myModal\").style.display = \"none\";\n      document.removeEventListener('keydown', closeModalOnEscape);\n    }\n\n    function closeModalOnEscape(event) {\n      if (event.key === \"Escape\") {\n        closeModal();\n      }\n    }\n\n    function closeModalOnClick(event) {\n      if (event.target === document.getElementById(\"myModal\")) {\n        closeModal();\n      }\n    }\n  </script></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func NewestBlogPost(posts []Post) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var14 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var14 == nil {
			templ_7745c5c3_Var14 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"newest-blog-post\" align=\"center\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 templ.SafeURL = templ.URL("/blog/" + strings.ToLower(posts[len(posts)-1].Title))
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var15)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs("📰 Latest Post - " + strings.ReplaceAll(posts[len(posts)-1].Title, "-", " ") + " 📰")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 502, Col: 93}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a><div class=\"scrolling-text-container\"><div class=\"scrolling-text\">&ensp;")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(posts[len(posts)-1].Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/blog.templ`, Line: 506, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
