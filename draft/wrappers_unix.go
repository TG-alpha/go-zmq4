// +build !windows

package zmq4

/*

#include <errno.h>
#include <zmq.h>

int zmq4_bind (void *socket, const char *endpoint)
{
    return zmq_bind(socket, endpoint);
}

int zmq4_close (void *socket)
{
    return zmq_close(socket);
}

int zmq4_connect (void *socket, const char *endpoint)
{
    return zmq_connect(socket, endpoint);
}

int zmq4_ctx_get (void *context, int option_name)
{
    return zmq_ctx_get(context, option_name);
}

void *zmq4_ctx_new ()
{
    return zmq_ctx_new();
}

int zmq4_ctx_set (void *context, int option_name, int option_value)
{
    return zmq_ctx_set(context, option_name, option_value);
}

int zmq4_ctx_term (void *context)
{
    return zmq_ctx_term(context);
}

int zmq4_curve_keypair (char *z85_public_key, char *z85_secret_key)
{
    return zmq_curve_keypair(z85_public_key, z85_secret_key);
}

int zmq4_curve_public (char *z85_public_key, char *z85_secret_key)
{
    return zmq_curve_public(z85_public_key, z85_secret_key);
}

int zmq4_disconnect (void *socket, const char *endpoint)
{
    return zmq_disconnect(socket, endpoint);
}

int zmq4_getsockopt (void *socket, int option_name, void *option_value, size_t *option_len)
{
    return zmq_getsockopt(socket, option_name, option_value, option_len);
}

const char *zmq4_msg_gets (zmq_msg_t *message, const char *property)
{
    return zmq_msg_gets(message, property);
}

int zmq4_msg_recv (zmq_msg_t *msg, void *socket, int flags)
{
    return zmq_msg_recv(msg, socket, flags);
}

int zmq4_poll (zmq_pollitem_t *items, int nitems, long timeout)
{
    return zmq_poll(items, nitems, timeout);
}

int zmq4_proxy (void *frontend, void *backend, void *capture)
{
    return zmq_proxy(frontend, backend, capture);
}

int zmq4_proxy_steerable (void *frontend, void *backend, void *capture, void *control)
{
    return zmq_proxy_steerable(frontend, backend, capture, control);
}

int zmq4_send (void *socket, void *buf, size_t len, int flags)
{
    return zmq_send(socket, buf, len, flags);
}

int zmq4_setsockopt (void *socket, int option_name, const void *option_value, size_t option_len)
{
    return zmq_setsockopt(socket, option_name, option_value, option_len);
}

void *zmq4_socket (void *context, int type)
{
    return zmq_socket(context, type);
}

int zmq4_socket_monitor (void *socket, char *endpoint, int events)
{
    return zmq_socket_monitor(socket, endpoint, events);
}

int zmq4_unbind (void *socket, const char *endpoint)
{
    return zmq_unbind(socket, endpoint);
}

// DRAFT

int zmq4_msg_init_size (zmq_msg_t *msg, size_t size)
{
    return zmq_msg_init_size(msg, size);
}

int zmq4_msg_send (zmq_msg_t *msg, void *socket, int flags)
{
    return zmq_msg_send(msg, socket, flags);
}

// not documented, so just assuming about return value...
int zmq4_msg_set_group (zmq_msg_t *msg, const char *group)
{
    return zmq_msg_set_group(msg, group);
}

int zmq4_msg_set_routing_id (zmq_msg_t *message, uint32_t routing_id)
{
    return zmq_msg_set_routing_id(message, routing_id);
}

// not documented, so just assuming about return value...
int zmq4_join (void *s, const char *group)
{
    return zmq4_join(s, group);
}

// not documented, so just assuming about return value...
int zmq4_leave (void *s, const char *group)
{
    return zmq4_leave(s, group);
}

*/
import "C"
