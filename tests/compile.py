import py_compile

files = ['simple.py']
for f in files:
    py_compile.compile(f, cfile=f+'c')
