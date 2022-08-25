def isValid(s):
    """
    :type s: str
    :rtype: bool
    """
    buf = []
    opbuf = {'(': ')', '[': ']', '{':'}'}
    
    for i in s:
        if i in opbuf:
            buf.append(opbuf[i])
        elif i in opbuf.values():
            if (not buf) or buf[-1] != i:
                return False
            else:
                buf.pop()
    if not buf:
        return True
    else:
        return False

s = "()"
print(isValid(s))