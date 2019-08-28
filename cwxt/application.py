import os
from logging import config
import tornado.ioloop
import sys
import yaml
import redis
import tornado.web
import tornado.gen
import tornado.httpserver
import tornado.options
from settings import settings_dict, database, user, password, host, port, CACHES
from tornado_mysql import pools
from handlers import urls
from backends.RedisCacheBackend import RedisCacheBackend


class App(tornado.web.Application):

    def __init__(self):
        handlers = urls
        redis_pool = redis.ConnectionPool(
            host=CACHES['HOST'],
            password=CACHES['PASSWORD'],
            port=CACHES['PORT']
        )
        self.cache = RedisCacheBackend(redis.Redis(connection_pool=redis_pool, charset="utf-8"))
        super(App, self).__init__(handlers=handlers, **settings_dict)
        self.pool = pools.Pool(dict(host=host, user=user, passwd=password, db=database, port=port, charset="utf8mb4"),
                               max_open_connections=1000, max_recycle_sec=50)


def options_define():
    tornado.options.define("port", default=8888, type=int, help="")


def main():
    options_define()
    log_config = yaml.safe_load(open("logconfig.yaml", "r"))
    config.dictConfig(log_config)
    tornado.options.parse_command_line()
    app = App()

    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(tornado.options.options.port)
    http_server.start()
    print(f"app run in port:{tornado.options.options.port}")
    tornado.ioloop.IOLoop.current().start()


if __name__ == '__main__':
    main()

































# /**             无可奉告 一颗赛艇
#  *  uJjYJYYLLv7r7vJJ5kqSFFFUUjJ7rrr7LLYLJLJ7
#  *  JuJujuYLrvuEM@@@B@@@B@B@B@@@MG5Y7vLjYjJL
#  *  JYjYJvr7XM@BB8GOOE8ZEEO8GqM8OBBBMu77LLJ7
#  *  LJLY7ru@@@BOZ8O8NXFFuSkSu25X0OFZ8MZJ;vLv
#  *  YvL7i5@BM8OGGqk22uvriiriii;r7LuSZXEMXrvr
#  *  vv7iU@BMNkF1uY7v7rr;iiii:i:i:ii7JEPNBPir
#  *  L7iL@BM8Xjuujvv77rr;ri;i;:iiiii:iLXFOBJ:
#  *  7ri@B@MOFuUS2Y7L7777rii;:::::i:iirjPG@O:
#  *  7:1B@BBOPjXXSJvrL7rr7iiii:i::::i;iv5MBB,
#  *  r:0@BBM8SFPX2Y77rri::iirri:::::iii75O@G.
#  *  7:SB@BBGqXPk0122UJL::i::r:::i:i;i:v2@Bk.
#  *  ri:MB@BBEqEMGq2JLLL1u7.iX51u77LF27iSB@r,
#  *  ri,v@B@MB8@qqNEqN1u:5B8BOFE0S7ii7qMB@F::
#  *  ii,J80Eq1MZkqPPX5YkPE@B@iXPE52j7:vBjE7::
#  *  ii:7MSqkS0PvLv7rrii0@L.Z1iLr::ir:rO,vi::
#  *  ii::EZXPSkquLvii:iF@N:.,BUi7ri,::UY;r:::
#  *  i::.2ONXqkPXS5FUUEOPP;..iSPXkjLYLLrr:::,
#  *  :::,iMXNP0NPLriiLGZ@BB1P87;JuL7r:7ri:::,
#  *  :::,.UGqNX0EZF2uUjUuULr:::,:7uuvv77::::.
#  *  ::::..5OXqXNJ50NSY;i:.,,,:i77Yvr;v;,,::.
#  *  :::,:.jOEPqPJiqBMMMO8NqP0SYLJriirv:.:,:.
#  *  ,:,,,.,Zq0P0X7vPFqF1ujLv7r:irrr7j7.,,::.
#  *  ,,,....0qk0080v75ujLLv7ri:i:rvj2J...,,,.
#  *  ......8@UXqZEMNvJjr;ii::,:::7uuv...,.,,.
#  *  .....B@BOvX88GMGk52vririiirJS1i.......,.
#  *  .JEMB@B@BMvL0MOMMMO8PE8GPqSk2L:.........
#  *  @B@@@B@M@B@L:7PGBOO8MOMOEP0Xri@B@Mk7,...
#  *  B@B@BBMBB@B@0::rJP8MO0uvvu7..,B@B@B@B@Z7
#  *  MMBM@BBB@B@B@Br:i,..:Lur:....7@OMMBM@B@@
#  *  8OOMMMOMMMMBB@B:....,PZENNi..JBOZ8GMOOOO
#  */
