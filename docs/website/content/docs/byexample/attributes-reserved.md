+++
# AUTOGENERATED BY byexample/generate.go
title= "Attributes-Reserved"
draft= false
description= ""
layout= "byexample"
weight = 7
topic = "Basics"
PlaygroundURL = "http://anz-bank.github.io/sysl-playground/?input=VXNlciBbfmh1bWFuXToKICAgIENoZWNrIEJhbGFuY2U6CiAgICAgICAgTW9iaWxlQXBwIDwtIExvZ2luCiAgICAgICAgTW9iaWxlQXBwIDwtIENoZWNrIEJhbGFuY2UKTW9iaWxlQXBwIFt+dWldOgogICAgTG9naW46CiAgICAgICAgU2VydmVyIDwtIExvZ2luCiAgICBDaGVjayBCYWxhbmNlOgogICAgICAgIFNlcnZlciA8LSBSZWFkIFVzZXIgQmFsYW5jZQpPYXV0aFNlcnZpY2VbfmV4dGVybmFsXToKICAgIE9hdXRoOiAuLi4KU2VydmVyOgogICAgTG9naW46CiAgICAgICAgZG8gaW5wdXQgdmFsaWRhdGlvbgogICAgICAgIE9hdXRoU2VydmljZSA8LSBPYXV0aAogICAgICAgIERCIDwtIFNhdmUKICAgICAgICByZXR1cm4gc3VjY2VzcyBvciBmYWlsdXJlCiAgICBSZWFkIFVzZXIgQmFsYW5jZToKICAgICAgICBEQiA8LSBMb2FkCiAgICAgICAgcmV0dXJuIGJhbGFuY2UKREIgW35kYl06CiAgICBTYXZlOiAuLi4KICAgIExvYWQ6IC4uLgpQcm9qZWN0IFtzZXF0aXRsZT0iRGlhZ3JhbSJdOgogICAgU2VxOgogICAgICAgIFVzZXIgPC0gQ2hlY2sgQmFsYW5jZQo=&cmd=c3lzbCBzZCAtbyAiM19wcm9qZWN0LnN2ZyIgLXMgIlByb2plY3QgPC0gU2VxIiAxX3Byb2plY3Quc3lzbA=="
ID = "attributes-reserved"
CodeWithoutComments = """User [~human]:
    Check Balance:
        MobileApp <- Login
        MobileApp <- Check Balance
MobileApp [~ui]:
    Login:
        Server <- Login
    Check Balance:
        Server <- Read User Balance
OauthService[~external]:
    Oauth: ...
Server:
    Login:
        do input validation
        OauthService <- Oauth
        DB <- Save
        return success or failure
    Read User Balance:
        DB <- Load
        return balance
DB [~db]:
    Save: ...
    Load: ...
Project [seqtitle="Diagram"]:
    Seq:
        User <- Check Balance
"""

Segs = [[
  
      {CodeEmpty= true,CodeLeading= true,CodeRun= false,CodeRendered="""""",DocsRendered= """<p>Sysl reserves certain attributes for sequence diagram generation. Reserved attributes alter the appearance of generated sequence diagrams</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= true,CodeRendered="""<pre class="chroma"><span class="nx">User</span> <span class="p">[</span><span class="err">~</span><span class="nx">human</span><span class="p">]:</span></pre>""",DocsRendered= """<p>~human specifies a user</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">    <span class="nx">Check</span> <span class="nx">Balance</span><span class="p">:</span>
        <span class="nx">MobileApp</span> <span class="o">&lt;-</span> <span class="nx">Login</span>
        <span class="nx">MobileApp</span> <span class="o">&lt;-</span> <span class="nx">Check</span> <span class="nx">Balance</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">MobileApp</span> <span class="p">[</span><span class="err">~</span><span class="nx">ui</span><span class="p">]:</span></pre>""",DocsRendered= """<p>~ui specifies a user interface</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">    <span class="nx">Login</span><span class="p">:</span>
        <span class="nx">Server</span> <span class="o">&lt;-</span> <span class="nx">Login</span>
    <span class="nx">Check</span> <span class="nx">Balance</span><span class="p">:</span>
        <span class="nx">Server</span> <span class="o">&lt;-</span> <span class="nx">Read</span> <span class="nx">User</span> <span class="nx">Balance</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">OauthService</span><span class="p">[</span><span class="err">~</span><span class="nx">external</span><span class="p">]:</span></pre>""",DocsRendered= """<p>~external specifies an external service</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">    <span class="nx">Oauth</span><span class="p">:</span> <span class="o">...</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">Server</span><span class="p">:</span>
    <span class="nx">Login</span><span class="p">:</span>
        <span class="nx">do</span> <span class="nx">input</span> <span class="nx">validation</span>
        <span class="nx">OauthService</span> <span class="o">&lt;-</span> <span class="nx">Oauth</span>
        <span class="nx">DB</span> <span class="o">&lt;-</span> <span class="nx">Save</span>
        <span class="k">return</span> <span class="nx">success</span> <span class="nx">or</span> <span class="nx">failure</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">    <span class="nx">Read</span> <span class="nx">User</span> <span class="nx">Balance</span><span class="p">:</span>
        <span class="nx">DB</span> <span class="o">&lt;-</span> <span class="nx">Load</span>
        <span class="k">return</span> <span class="nx">balance</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">DB</span> <span class="p">[</span><span class="err">~</span><span class="nx">db</span><span class="p">]:</span></pre>""",DocsRendered= """<p>~db specifies a database</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= true,CodeRun= false,CodeRendered="""<pre class="chroma">    <span class="nx">Save</span><span class="p">:</span> <span class="o">...</span>
    <span class="nx">Load</span><span class="p">:</span> <span class="o">...</span></pre>""",DocsRendered= """""",Image = ""},

      {CodeEmpty= false,CodeLeading= false,CodeRun= false,CodeRendered="""<pre class="chroma"><span class="nx">Project</span> <span class="p">[</span><span class="nx">seqtitle</span><span class="p">=</span><span class="s">&#34;Diagram&#34;</span><span class="p">]:</span>
    <span class="nx">Seq</span><span class="p">:</span>
        <span class="nx">User</span> <span class="o">&lt;-</span> <span class="nx">Check</span> <span class="nx">Balance</span></pre>""",DocsRendered= """""",Image = ""},


],
[
  
      {CodeEmpty= true,CodeLeading= true,CodeRun= false,CodeRendered="""""",DocsRendered= """<p>We can see the effect of the reserved attributes on the generated diagram</p>
""",Image = ""},

      {CodeEmpty= false,CodeLeading= false,CodeRun= true,CodeRendered="""<pre class="chroma"><span class="nx">sysl</span> <span class="nx">sd</span> <span class="o">-</span><span class="nx">o</span> <span class="s">&#34;3_project.svg&#34;</span> <span class="o">-</span><span class="nx">s</span> <span class="s">&#34;Project &lt;- Seq&#34;</span> <span class="mi">1</span><span class="nx">_project</span><span class="p">.</span><span class="nx">sysl</span></pre>""",DocsRendered= """""",Image = ""},


],
[
  
      {CodeEmpty= false,CodeLeading= false,CodeRun= false,CodeRendered="""""",DocsRendered= """""",Image = "/assets/byexample/images/attributes-reserved6.svg"},


],

]
+++


