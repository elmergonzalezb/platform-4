from functools import update_wrapper
from flask import make_response, redirect, request
from syncloud_platform.injector import get_injector


def nocache(f):
    def new_func(*args, **kwargs):
        resp = make_response(f(*args, **kwargs))
        # resp.cache_control.no_cache = True
        resp.headers['Cache-Control'] = 'no-store, no-cache, must-revalidate, post-check=0, pre-check=0, max-age=0'
        return resp
    return update_wrapper(new_func, f)


def redirect_if_not_activated(f):
    platform_user_config = get_injector().user_platform_config

    def new_func(*args, **kwargs):
        resp = make_response(f(*args, **kwargs))
        if not platform_user_config.is_activated():
            return redirect('http://{0}:81'.format(request.host))
        else:
            return resp

    return update_wrapper(new_func, f)
