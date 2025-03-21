# sont décrits ici les lexèmes de MarkDown

# ============= leaf blocks ============

class THEME_BREAK():

    def __init__(self, content):
        self.content = content
        self.one_line = True

class TITLE():

    def __init__(self, content):
        self.content = content
        self.one_line = True

class IDENTED_CODE_OPN():

    def __init__(self, content):
        self.content = content
        self.one_line = False

class INDENTED_CODE_CL():

    def __init__(self, content):
        self.content = content
        self.one_line = False

class FENCED_CODE_OPN():

    def __init__(self, content, language = None):
        self.content = content
        self.language = language
        self.one_line = False

class HTML_BLOCK():
    pass

    #TODO : un peu complexe pour l'instant


class PARAGRAPH_OPN():

    def __init__(self, content):
        self.content = content
        self.one_line = False

class PARAGRAPH_CL():

    def __init__(self, content):
        self.content = content
        self.one_line = False

class BLANK_LINE():

    def __init__(self, content):
        self.content = content
        self.one_line = True

class TABLE():
    pass 

    #TODO : un peu complexe pour l'instant