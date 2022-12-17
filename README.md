# MicroGoogle
A google scraping RESTful API.

# Endpoints & Usage

To search google, construct a JSON Query with a structure similar to:
```
{
    "pages": "1",
    "query": "minecraft"
}
```
and POST it to `/search/google`. If nothing goes wrong, you should get a response that resembles:

```
{
    "response": [
        {
            "breadcrumb": "https://www.minecraft.net",
            "description": "Explore new gaming adventures, accessories, & merchandise on the Minecraft Official Site. Buy & download the game here, or check the site for the latest ...",
            "link": "https://www.minecraft.net/",
            "title": "Welcome to the Minecraft Official Site | Minecraft"
        },
        ...
    ]
}
```

# Dependencies
https://github.com/openinfolabs-org/serp -> The google scraper, slightly modified.

# License
This code is license under the MIT license.
Copyright 2022 SteveGremory

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

