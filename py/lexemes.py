# sont décrits ici les lexèmes de MarkDown

# ============= leaf blocks ============

class THEME_BREAK():
    one_line = True

    def __init__(self, content):
        self.content = content

class TITLE():
    one_line = True

    def __init__(self, content):
        self.content = content

class IDENTED_CODE_OPN():
    one_line = False

    def __init__(self, content):
        self.content = content

class INDENTED_CODE_CL():
    one_line = False

    def __init__(self, content):
        self.content = content

class FENCED_CODE_OPN():
    one_line = False

    def __init__(self, content, language = None):
        self.content = content
        self.language = language

class HTML_BLOCK():
    pass

    #TODO : un peu complexe pour l'instant


class PARAGRAPH_OPN():
    one_line = False

    def __init__(self, content):
        self.content = content

class PARAGRAPH_CL():
    one_line = False

    def __init__(self, content):
        self.content = content

class BLANK_LINE():
    one_line = True

    def __init__(self, content):
        self.content = content

class TABLE():
    pass 

    #TODO : un peu complexe pour l'instant