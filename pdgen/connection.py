from pdgen import config
from tabulate import tabulate


def __name(name: str):
    if name is None:
        name = "DEFAULT"

    return name


def __check_name(name: str):
    msg = ""
    
    cfg = config.read()
    if cfg == None: 
        return "There are no connection defined."
    
    conns = [x for x in cfg.keys()]
  
        

    if name not in conns:
        msg = "There's no connection named '{}'.\n".format(name)
        msg += "Use 'pdgen connection list' to list all connections."

    return msg


def __permission_error(operation: str):
    msg = 'Unable to {} a connection with URL and Name provided. '
    msg += 'It looks like a permission problem at the .pdgen file.\n{}'

    return msg.format(operation, config.config_file)


def add(url: str, name: str) -> str:
    name = __name(name=name)

    if not config.check_connection(url):
        msg = 'Unable to connect using the provided URL.\n'
        msg += 'url={}'.format(url)

        return msg

    if config.add_connection(url, name) is False:
        return __permission_error("add")

    return 'Connection created successfully.\n(name:{}, url:{})'.format(name, url)


def rm(name: str) -> str:
    cfg = config.read()
    if cfg is None: return __permission_error('remove')
    
    cn = __check_name(name=name.upper())
    if len(cn) > 0: return cn

    if config.remove_connection(name) is False:
        return 'Sorry, some unknown error happened'

    msg = 'Connection removed successfully.\n'
    msg += "Use 'pdgen connection list' to list all connections."
    
    return msg


def lst() -> str:
    conn_list = config.list_connections()

    if len(conn_list) == 0:
        return "There is no added connection."

    table = tabulate(
        conn_list,
        headers=['Name', 'URL', 'Selected'],
        tablefmt='orgtbl'
    )
    
    return '\n' + str(table) + '\n'


def use(name: str) -> str:
    name = name.upper()

    cn = __check_name(name=name)
    if len(cn) > 0:
        return cn

    if config.select_connection(name=name) is False:
        return __permission_error("add")

    return "The connection named '{}' was selected to use.".format(name)


