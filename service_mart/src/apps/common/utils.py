import random
import string


def generate_bsc(length):
    """
    生成随机码（length位数字与字母混合大小写）
    :return bsr:length 位数字与字母混合大小写
    """
    bsc = ''.join(random.choice(string.ascii_letters + string.digits) for _ in range(length))
    return bsc