import base64
from hashlib import sha256

reqDetails = """GET
/v1/user/orders?open=true
1573604427

vwj043jtrw4o5igw4oi5jwoi45g"""

if __name__ == "__main__":
    hash = sha256(reqDetails.encode("utf8")).digest()
    signature = base64.b64encode(hash)
    print(signature)