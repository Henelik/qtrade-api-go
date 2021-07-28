from qtrade_client.api import QtradeAPI, APIException
from pprint import pprint

if __name__ == "__main__":
    api = QtradeAPI("https://api.qtrade.io", key="")

    for code, currency in api.currencies.items():
        print(str(code) + ": " + str(currency["precision"]))