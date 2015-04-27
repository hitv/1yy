(function(a) {
    if (String.prototype.trim === a) {
        String.prototype.trim = function() {
            return this.replace(/^\s+/, "").replace(/\s+$/, "");
        };
    }
    if (Array.prototype.reduce === a) {
        Array.prototype.reduce = function(c) {
            if (this === void 0 || this === null) {
                throw new TypeError();
            }
            var f = Object(this), b = f.length >>> 0, e = 0, d;
            if (typeof c != "function") {
                throw new TypeError();
            }
            if (b == 0 && arguments.length == 1) {
                throw new TypeError();
            }
            if (arguments.length >= 2) {
                d = arguments[1];
            } else {
                do {
                    if (e in f) {
                        d = f[e++];
                        break;
                    }
                    if (++e >= b) {
                        throw new TypeError();
                    }
                } while (true);
            }
            while (e < b) {
                if (e in f) {
                    d = c.call(a, d, f[e], e, f);
                }
                e++;
            }
            return d;
        };
    }
})();
var Zepto = (function() {
    var m, t, E, a, M = [], o = M.slice, F = M.filter, h = window.document, L = {}, N = {}, q = h.defaultView.getComputedStyle, V = {"column-count": 1,columns: 1,"font-weight": 1,"line-height": 1,opacity: 1,"z-index": 1,zoom: 1}, x = /^\s*<(\w+|!)[^>]*>/, k = /<(?!area|br|col|embed|hr|img|input|link|meta|param)(([\w:]+)[^>]*)\/>/ig, I = /^(?:body|html)$/i, D = ["val", "css", "html", "text", "data", "width", "height", "offset"], y = ["after", "prepend", "before", "append"], u = h.createElement("table"), O = h.createElement("tr"), i = {tr: h.createElement("tbody"),tbody: u,thead: u,tfoot: u,td: O,th: O,"*": h.createElement("div")}, v = /complete|loaded|interactive/, H = /^\.([\w-]+)$/, w = /^#([\w-]*)$/, K = /^[\w-]+$/, e = {}, g = e.toString, d = {}, T, P, G = h.createElement("div");
    d.matches = function(ac, Y) {
        if (!ac || ac.nodeType !== 1) {
            return false;
        }
        var aa = ac.webkitMatchesSelector || ac.mozMatchesSelector || ac.oMatchesSelector || ac.matchesSelector;
        if (aa) {
            return aa.call(ac, Y);
        }
        var ab, ad = ac.parentNode, Z = !ad;
        if (Z) {
            (ad = G).appendChild(ac);
        }
        ab = ~d.qsa(ad, Y).indexOf(ac);
        Z && G.removeChild(ac);
        return ab;
    };
    function X(Y) {
        return Y == null ? String(Y) : e[g.call(Y)] || "object";
    }
    function p(Y) {
        return g.call(Y) == "[object Function]";
    }
    function J(Y) {
        return Y instanceof Object;
    }
    function W(Y) {
        return J(Y) && Y.__proto__ == Object.prototype;
    }
    function A(Y) {
        return Y instanceof Array;
    }
    function B(Y) {
        return typeof Y.length == "number";
    }
    function U(Y) {
        return F.call(Y, function(Z) {
            return Z !== m && Z !== null;
        });
    }
    function C(Y) {
        return Y.length > 0 ? E.fn.concat.apply([], Y) : Y;
    }
    T = function(Y) {
        return Y.replace(/-+(.)?/g, function(Z, aa) {
            return aa ? aa.toUpperCase() : "";
        });
    };
    function n(Y) {
        return Y.replace(/::/g, "/").replace(/([A-Z]+)([A-Z][a-z])/g, "$1_$2").replace(/([a-z\d])([A-Z])/g, "$1_$2").replace(/_/g, "-").toLowerCase();
    }
    P = function(Y) {
        return F.call(Y, function(aa, Z) {
            return Y.indexOf(aa) == Z;
        });
    };
    function Q(Y) {
        return Y in N ? N[Y] : (N[Y] = new RegExp("(^|\\s)" + Y + "(\\s|$)"));
    }
    function f(Y, Z) {
        return (typeof Z == "number" && !V[n(Y)]) ? Z + "px" : Z;
    }
    function R(aa) {
        var Y, Z;
        if (!L[aa]) {
            Y = h.createElement(aa);
            h.body.appendChild(Y);
            Z = q(Y, "").getPropertyValue("display");
            Y.parentNode.removeChild(Y);
            Z == "none" && (Z = "block");
            L[aa] = Z;
        }
        return L[aa];
    }
    function s(Y) {
        return "children" in Y ? o.call(Y.children) : E.map(Y.childNodes, function(Z) {
            if (Z.nodeType == 1) {
                return Z;
            }
        });
    }
    d.fragment = function(ac, aa, ab) {
        if (ac.replace) {
            ac = ac.replace(k, "<$1></$2>");
        }
        if (aa === m) {
            aa = x.test(ac) && RegExp.$1;
        }
        if (!(aa in i)) {
            aa = "*";
        }
        var Z, ad, Y = i[aa];
        Y.innerHTML = "" + ac;
        ad = E.each(o.call(Y.childNodes), function() {
            Y.removeChild(this);
        });
        if (W(ab)) {
            Z = E(ad);
            E.each(ab, function(ae, af) {
                if (D.indexOf(ae) > -1) {
                    Z[ae](af);
                } else {
                    Z.attr(ae, af);
                }
            });
        }
        return ad;
    };
    d.Z = function(Z, Y) {
        Z = Z || [];
        Z.__proto__ = arguments.callee.prototype;
        Z.selector = Y || "";
        return Z;
    };
    d.isZ = function(Y) {
        return Y instanceof d.Z;
    };
    d.init = function(Y, Z) {
        if (!Y) {
            return d.Z();
        } else {
            if (p(Y)) {
                return E(h).ready(Y);
            } else {
                if (d.isZ(Y)) {
                    return Y;
                } else {
                    var aa;
                    if (A(Y)) {
                        aa = U(Y);
                    } else {
                        if (J(Y)) {
                            aa = [W(Y) ? E.extend({}, Y) : Y], Y = null;
                        } else {
                            if (x.test(Y)) {
                                aa = d.fragment(Y.trim(), RegExp.$1, Z), Y = null;
                            } else {
                                if (Z !== m) {
                                    return E(Z).find(Y);
                                } else {
                                    aa = d.qsa(h, Y);
                                }
                            }
                        }
                    }
                    return d.Z(aa, Y);
                }
            }
        }
    };
    E = function(Y, Z) {
        return d.init(Y, Z);
    };
    function l(aa, Z, Y) {
        for (t in Z) {
            if (Y && W(Z[t])) {
                if (!W(aa[t])) {
                    aa[t] = {};
                }
                l(aa[t], Z[t], Y);
            } else {
                if (Z[t] !== m) {
                    aa[t] = Z[t];
                }
            }
        }
    }
    E.extend = function(aa) {
        var Y, Z = o.call(arguments, 1);
        if (typeof aa == "boolean") {
            Y = aa;
            aa = Z.shift();
        }
        Z.forEach(function(ab) {
            l(aa, ab, Y);
        });
        return aa;
    };
    d.qsa = function(Z, Y) {
        var aa;
        return (Z === h && w.test(Y)) ? ((aa = Z.getElementById(RegExp.$1)) ? [aa] : []) : (Z.nodeType !== 1 && Z.nodeType !== 9) ? [] : o.call(H.test(Y) ? Z.getElementsByClassName(RegExp.$1) : K.test(Y) ? Z.getElementsByTagName(Y) : Z.querySelectorAll(Y));
    };
    function z(Z, Y) {
        return Y === m ? E(Z) : E(Z).filter(Y);
    }
    E.contains = function(Y, Z) {
        return Y !== Z && Y.contains(Z);
    };
    function r(aa, Z, Y, ab) {
        return p(Z) ? Z.call(aa, Y, ab) : Z;
    }
    function b(Z, Y, aa) {
        aa == null ? Z.removeAttribute(Y) : Z.setAttribute(Y, aa);
    }
    function S(aa, ab) {
        var Y = aa.className, Z = Y && Y.baseVal !== m;
        if (ab === m) {
            return Z ? Y.baseVal : Y;
        }
        Z ? (Y.baseVal = ab) : (aa.className = ab);
    }
    function j(Z) {
        var Y;
        try {
            return Z ? Z == "true" || (Z == "false" ? false : Z == "null" ? null : !isNaN(Y = Number(Z)) ? Y : /^[\[\{]/.test(Z) ? E.parseJSON(Z) : Z) : Z;
        } catch (aa) {
            return Z;
        }
    }
    E.type = X;
    E.isFunction = p;
    E.isObject = J;
    E.isArray = A;
    E.isPlainObject = W;
    E.inArray = function(Z, aa, Y) {
        return M.indexOf.call(aa, Z, Y);
    };
    E.camelCase = d.camelize = T;
    E.trim = function(Y) {
        return Y.trim();
    };
    E.uuid = 0;
    E.support = {};
    E.expr = {};
    E.map = function(ac, ad) {
        var ab, Y = [], aa, Z;
        if (B(ac)) {
            for (aa = 0; aa < ac.length; aa++) {
                ab = ad(ac[aa], aa);
                if (ab != null) {
                    Y.push(ab);
                }
            }
        } else {
            for (Z in ac) {
                ab = ad(ac[Z], Z);
                if (ab != null) {
                    Y.push(ab);
                }
            }
        }
        return C(Y);
    };
    E.each = function(aa, ab) {
        var Z, Y;
        if (B(aa)) {
            for (Z = 0; Z < aa.length; Z++) {
                if (ab.call(aa[Z], Z, aa[Z]) === false) {
                    return aa;
                }
            }
        } else {
            for (Y in aa) {
                if (ab.call(aa[Y], Y, aa[Y]) === false) {
                    return aa;
                }
            }
        }
        return aa;
    };
    E.grep = function(Y, Z) {
        return F.call(Y, Z);
    };
    if (window.JSON) {
        E.parseJSON = JSON.parse;
    }
    E.each("Boolean Number String Function Array Date RegExp Object Error".split(" "), function(Z, Y) {
        e["[object " + Y + "]"] = Y.toLowerCase();
    });
    E.fn = {forEach: M.forEach,reduce: M.reduce,push: M.push,sort: M.sort,indexOf: M.indexOf,concat: M.concat,map: function(Y) {
            return E(E.map(this, function(aa, Z) {
                return Y.call(aa, Z, aa);
            }));
        },slice: function() {
            return E(o.apply(this, arguments));
        },ready: function(Y) {
            if (v.test(h.readyState)) {
                Y(E);
            } else {
                h.addEventListener("DOMContentLoaded", function() {
                    Y(E);
                }, false);
            }
            return this;
        },get: function(Y) {
            return Y === m ? o.call(this) : this[Y];
        },toArray: function() {
            return this.get();
        },size: function() {
            return this.length;
        },remove: function() {
            return this.each(function() {
                if (this.parentNode != null) {
                    this.parentNode.removeChild(this);
                }
            });
        },each: function(Y) {
            this.forEach(function(aa, Z) {
                Y.call(aa, Z, aa);
            });
            return this;
        },filter: function(Y) {
            if (p(Y)) {
                return this.not(this.not(Y));
            }
            return E(F.call(this, function(Z) {
                return d.matches(Z, Y);
            }));
        },add: function(Y, Z) {
            return E(P(this.concat(E(Y, Z))));
        },is: function(Y) {
            return this.length > 0 && d.matches(this[0], Y);
        },not: function(Y) {
            var Z = [];
            if (p(Y) && Y.call !== m) {
                this.each(function(ab) {
                    if (!Y.call(this, ab)) {
                        Z.push(this);
                    }
                });
            } else {
                var aa = typeof Y == "string" ? this.filter(Y) : (B(Y) && p(Y.item)) ? o.call(Y) : E(Y);
                this.forEach(function(ab) {
                    if (aa.indexOf(ab) < 0) {
                        Z.push(ab);
                    }
                });
            }
            return E(Z);
        },has: function(Y) {
            return this.filter(function() {
                return J(Y) ? E.contains(this, Y) : E(this).find(Y).size();
            });
        },eq: function(Y) {
            return Y === -1 ? this.slice(Y) : this.slice(Y, +Y + 1);
        },first: function() {
            var Y = this[0];
            return Y && !J(Y) ? Y : E(Y);
        },last: function() {
            var Y = this[this.length - 1];
            return Y && !J(Y) ? Y : E(Y);
        },find: function(Z) {
            var Y;
            if (this.length == 1) {
                Y = E(d.qsa(this[0], Z));
            } else {
                Y = this.map(function() {
                    return d.qsa(this, Z);
                });
            }
            return Y;
        },closest: function(Y, Z) {
            var aa = this[0];
            while (aa && !d.matches(aa, Y)) {
                aa = aa !== Z && aa !== h && aa.parentNode;
            }
            return E(aa);
        },parents: function(Y) {
            var aa = [], Z = this;
            while (Z.length > 0) {
                Z = E.map(Z, function(ab) {
                    if ((ab = ab.parentNode) && ab !== h && aa.indexOf(ab) < 0) {
                        aa.push(ab);
                        return ab;
                    }
                });
            }
            return z(aa, Y);
        },parent: function(Y) {
            return z(P(this.pluck("parentNode")), Y);
        },children: function(Y) {
            return z(this.map(function() {
                return s(this);
            }), Y);
        },contents: function() {
            return this.map(function() {
                return o.call(this.childNodes);
            });
        },siblings: function(Y) {
            return z(this.map(function(Z, aa) {
                return F.call(s(aa.parentNode), function(ab) {
                    return ab !== aa;
                });
            }), Y);
        },empty: function() {
            return this.each(function() {
                this.innerHTML = "";
            });
        },pluck: function(Y) {
            return E.map(this, function(Z) {
                return Z[Y];
            });
        },show: function() {
            return this.each(function() {
                this.style.display == "none" && (this.style.display = null);
                if (q(this, "").getPropertyValue("display") == "none") {
                    this.style.display = R(this.nodeName);
                }
            });
        },replaceWith: function(Y) {
            return this.before(Y).remove();
        },wrap: function(Y) {
            var Z = p(Y);
            if (this[0] && !Z) {
                var aa = E(Y).get(0), ab = aa.parentNode || this.length > 1;
            }
            return this.each(function(ac) {
                E(this).wrapAll(Z ? Y.call(this, ac) : ab ? aa.cloneNode(true) : aa);
            });
        },wrapAll: function(Y) {
            if (this[0]) {
                E(this[0]).before(Y = E(Y));
                var Z;
                while ((Z = Y.children()).length) {
                    Y = Z.first();
                }
                E(Y).append(this);
            }
            return this;
        },wrapInner: function(Y) {
            var Z = p(Y);
            return this.each(function(ab) {
                var aa = E(this), ac = aa.contents(), ad = Z ? Y.call(this, ab) : Y;
                ac.length ? ac.wrapAll(ad) : aa.append(ad);
            });
        },unwrap: function() {
            this.parent().each(function() {
                E(this).replaceWith(E(this).children());
            });
            return this;
        },clone: function() {
            return this.map(function() {
                return this.cloneNode(true);
            });
        },hide: function() {
            return this.css("display", "none");
        },toggle: function(Y) {
            return this.each(function() {
                var Z = E(this);
                (Y === m ? Z.css("display") == "none" : Y) ? Z.show() : Z.hide();
            });
        },prev: function(Y) {
            return E(this.pluck("previousElementSibling")).filter(Y || "*");
        },next: function(Y) {
            return E(this.pluck("nextElementSibling")).filter(Y || "*");
        },html: function(Y) {
            return Y === m ? (this.length > 0 ? this[0].innerHTML : null) : this.each(function(Z) {
                var aa = this.innerHTML;
                E(this).empty().append(r(this, Y, Z, aa));
            });
        },text: function(Y) {
            return Y === m ? (this.length > 0 ? this[0].textContent : null) : this.each(function() {
                this.textContent = Y;
            });
        },attr: function(Z, aa) {
            var Y;
            return (typeof Z == "string" && aa === m) ? (this.length == 0 || this[0].nodeType !== 1 ? m : (Z == "value" && this[0].nodeName == "INPUT") ? this.val() : (!(Y = this[0].getAttribute(Z)) && Z in this[0]) ? this[0][Z] : Y) : this.each(function(ab) {
                if (this.nodeType !== 1) {
                    return;
                }
                if (J(Z)) {
                    for (t in Z) {
                        b(this, t, Z[t]);
                    }
                } else {
                    b(this, Z, r(this, aa, ab, this.getAttribute(Z)));
                }
            });
        },removeAttr: function(Y) {
            return this.each(function() {
                this.nodeType === 1 && b(this, Y);
            });
        },prop: function(Y, Z) {
            return (Z === m) ? (this[0] ? this[0][Y] : m) : this.each(function(aa) {
                this[Y] = r(this, Z, aa, this[Y]);
            });
        },data: function(Y, aa) {
            var Z = this.attr("data-" + n(Y), aa);
            return Z !== null ? Z : m;
        },val: function(Y) {
            return (Y === m) ? (this.length > 0 ? (this[0].multiple ? E(this[0]).find("option").filter(function(Z) {
                return this.selected;
            }).pluck("value") : this[0].value) : m) : this.each(function(Z) {
                this.value = r(this, Y, Z, this.value);
            });
        },offset: function() {
            if (this.length == 0) {
                return null;
            }
            var Y = this[0].getBoundingClientRect();
            return {left: Y.left + window.pageXOffset,top: Y.top + window.pageYOffset,width: Y.width,height: Y.height};
        },css: function(aa, Z) {
            if (arguments.length < 2 && typeof aa == "string") {
                return (this.length == 0 ? m : this[0].style[T(aa)] || q(this[0], "").getPropertyValue(aa));
            }
            var Y = "";
            for (t in aa) {
                if (!aa[t] && aa[t] !== 0) {
                    this.each(function() {
                        this.style.removeProperty(n(t));
                    });
                } else {
                    Y += n(t) + ":" + f(t, aa[t]) + ";";
                }
            }
            if (typeof aa == "string") {
                if (!Z && Z !== 0) {
                    this.each(function() {
                        this.style.removeProperty(n(aa));
                    });
                } else {
                    Y = n(aa) + ":" + f(aa, Z);
                }
            }
            return this.each(function() {
                this.style.cssText += ";" + Y;
            });
        },index: function(Y) {
            return Y ? this.indexOf(E(Y)[0]) : this.parent().children().indexOf(this[0]);
        },hasClass: function(Y) {
            if (this.length < 1) {
                return false;
            } else {
                return Q(Y).test(S(this[0]));
            }
        },addClass: function(Y) {
            return this.each(function(Z) {
                a = [];
                var ab = S(this), aa = r(this, Y, Z, ab);
                aa.split(/\s+/g).forEach(function(ac) {
                    if (!E(this).hasClass(ac)) {
                        a.push(ac);
                    }
                }, this);
                a.length && S(this, ab + (ab ? " " : "") + a.join(" "));
            });
        },removeClass: function(Y) {
            return this.each(function(Z) {
                if (Y === m) {
                    return S(this, "");
                }
                a = S(this);
                r(this, Y, Z, a).split(/\s+/g).forEach(function(aa) {
                    a = a.replace(Q(aa), " ");
                });
                S(this, a.trim());
            });
        },toggleClass: function(Z, Y) {
            return this.each(function(aa) {
                var ab = r(this, Z, aa, S(this));
                (Y === m ? !E(this).hasClass(ab) : Y) ? E(this).addClass(ab) : E(this).removeClass(ab);
            });
        },scrollTop: function() {
            if (!this.length) {
                return;
            }
            return ("scrollTop" in this[0]) ? this[0].scrollTop : this[0].scrollY;
        },position: function() {
            if (!this.length) {
                return;
            }
            var aa = this[0], Z = this.offsetParent(), ab = this.offset(), Y = I.test(Z[0].nodeName) ? {top: 0,left: 0} : Z.offset();
            ab.top -= parseFloat(E(aa).css("margin-top")) || 0;
            ab.left -= parseFloat(E(aa).css("margin-left")) || 0;
            Y.top += parseFloat(E(Z[0]).css("border-top-width")) || 0;
            Y.left += parseFloat(E(Z[0]).css("border-left-width")) || 0;
            return {top: ab.top - Y.top,left: ab.left - Y.left};
        },offsetParent: function() {
            return this.map(function() {
                var Y = this.offsetParent || h.body;
                while (Y && !I.test(Y.nodeName) && E(Y).css("position") == "static") {
                    Y = Y.offsetParent;
                }
                return Y;
            });
        }};
    ["width", "height"].forEach(function(Y) {
        E.fn[Y] = function(Z) {
            var ab, aa = Y.replace(/./, function(ac) {
                return ac[0].toUpperCase();
            });
            if (Z === m) {
                return this[0] == window ? window["inner" + aa] : this[0] == h ? h.documentElement["offset" + aa] : (ab = this.offset()) && ab[Y];
            } else {
                return this.each(function(ac) {
                    var ad = E(this);
                    ad.css(Y, r(this, Z, ac, ad[Y]()));
                });
            }
        };
    });
    function c(aa, Y) {
        Y(aa);
        for (var Z in aa.childNodes) {
            c(aa.childNodes[Z], Y);
        }
    }
    y.forEach(function(aa, Z) {
        var Y = Z % 2;
        E.fn[aa] = function() {
            var ab = E.map(arguments, function(ae) {
                return J(ae) ? ae : d.fragment(ae);
            }), ac, ad = this.length > 1;
            if (ab.length < 1) {
                return this;
            }
            return this.each(function(ae, af) {
                ac = Y ? af : af.parentNode;
                af = Z == 0 ? af.nextSibling : Z == 1 ? af.firstChild : Z == 2 ? af : null;
                ab.forEach(function(ag) {
                    if (ad) {
                        ag = ag.cloneNode(true);
                    } else {
                        if (!ac) {
                            return E(ag).remove();
                        }
                    }
                    c(ac.insertBefore(ag, af), function(ah) {
                        if (ah.nodeName != null && ah.nodeName.toUpperCase() === "SCRIPT" && (!ah.type || ah.type === "text/javascript") && !ah.src) {
                            window["eval"].call(window, ah.innerHTML);
                        }
                    });
                });
            });
        };
        E.fn[Y ? aa + "To" : "insert" + (Z ? "Before" : "After")] = function(ab) {
            E(ab)[aa](this);
            return this;
        };
    });
    d.Z.prototype = E.fn;
    d.uniq = P;
    d.deserializeValue = j;
    E.zepto = d;
    return E;
})();
window.Zepto = Zepto;
"$" in window || (window.$ = Zepto);
// ajax
(function($) {
    var jsonpID = 0, document = window.document, key, name, rscript = /<script\b[^<]*(?:(?!<\/script>)<[^<]*)*<\/script>/gi, scriptTypeRE = /^(?:text|application)\/javascript/i, xmlTypeRE = /^(?:text|application)\/xml/i, jsonType = "application/json", htmlType = "text/html", blankRE = /^\s*$/;
    function triggerAndReturn(context, eventName, data) {
        var event = $.Event(eventName);
        $(context).trigger(event, data);
        return !event.defaultPrevented;
    }
    function triggerGlobal(settings, context, eventName, data) {
        if (settings.global) {
            return triggerAndReturn(context || document, eventName, data);
        }
    }
    $.active = 0;
    function ajaxStart(settings) {
        if (settings.global && $.active++ === 0) {
            triggerGlobal(settings, null, "ajaxStart");
        }
    }
    function ajaxStop(settings) {
        if (settings.global && !(--$.active)) {
            triggerGlobal(settings, null, "ajaxStop");
        }
    }
    function ajaxBeforeSend(xhr, settings) {
        var context = settings.context;
        if (settings.beforeSend.call(context, xhr, settings) === false || triggerGlobal(settings, context, "ajaxBeforeSend", [xhr, settings]) === false) {
            return false;
        }
        triggerGlobal(settings, context, "ajaxSend", [xhr, settings]);
    }
    function ajaxSuccess(data, xhr, settings) {
        var context = settings.context, status = "success";
        settings.success.call(context, data, status, xhr);
        triggerGlobal(settings, context, "ajaxSuccess", [xhr, settings, data]);
        ajaxComplete(status, xhr, settings);
    }
    function ajaxError(error, type, xhr, settings) {
        var context = settings.context;
        settings.error.call(context, xhr, type, error);
        triggerGlobal(settings, context, "ajaxError", [xhr, settings, error]);
        ajaxComplete(type, xhr, settings);
    }
    function ajaxComplete(status, xhr, settings) {
        var context = settings.context;
        settings.complete.call(context, xhr, status);
        triggerGlobal(settings, context, "ajaxComplete", [xhr, settings]);
        ajaxStop(settings);
    }
    function empty() {
    }
    $.ajaxJSONP = function(options) {
        if (!("type" in options)) {
            return $.ajax(options);
        }
        var callbackName = "jsonp" + (++jsonpID), script = document.createElement("script"), cleanup = function() {
            clearTimeout(abortTimeout);
            $(script).remove();
            delete window[callbackName];
        }, abort = function(type) {
            cleanup();
            if (!type || type == "timeout") {
                window[callbackName] = empty;
            }
            ajaxError(null, type || "abort", xhr, options);
        }, xhr = {abort: abort}, abortTimeout;
        if (ajaxBeforeSend(xhr, options) === false) {
            abort("abort");
            return false;
        }
        window[callbackName] = function(data) {
            cleanup();
            ajaxSuccess(data, xhr, options);
        };
        script.onerror = function() {
            abort("error");
        };
        script.src = options.url.replace(/=\?/, "=" + callbackName);
        $("head").append(script);
        if (options.timeout > 0) {
            abortTimeout = setTimeout(function() {
                abort("timeout");
            }, options.timeout);
        }
        return xhr;
    };
    $.ajaxSettings = {type: "GET",beforeSend: empty,success: empty,error: empty,complete: empty,context: null,global: true,xhr: function() {
            return new window.XMLHttpRequest();
        },accepts: {script: "text/javascript, application/javascript",json: jsonType,xml: "application/xml, text/xml",html: htmlType,text: "text/plain"},crossDomain: false,timeout: 0,processData: true,cache: false,};
    function mimeToDataType(mime) {
        if (mime) {
            mime = mime.split(";", 2)[0];
        }
        return mime && (mime == htmlType ? "html" : mime == jsonType ? "json" : scriptTypeRE.test(mime) ? "script" : xmlTypeRE.test(mime) && "xml") || "text";
    }
    function appendQuery(url, query) {
        return (url + "&" + query).replace(/[&?]{1,2}/, "?");
    }
    function serializeData(options) {
        if (options.processData && options.data && $.type(options.data) != "string") {
            options.data = $.param(options.data, options.traditional);
        }
        if (options.data && (!options.type || options.type.toUpperCase() == "GET")) {
            options.url = appendQuery(options.url, options.data);
        }
    }
    $.ajax = function(options) {
        var settings = $.extend({}, options || {});
        for (key in $.ajaxSettings) {
            if (settings[key] === undefined) {
                settings[key] = $.ajaxSettings[key];
            }
        }
        ajaxStart(settings);
        if (!settings.crossDomain) {
            settings.crossDomain = /^([\w-]+:)?\/\/([^\/]+)/.test(settings.url) && RegExp.$2 != window.location.host;
        }
        if (!settings.url) {
            settings.url = window.location.toString();
        }
        serializeData(settings);
        if (settings.cache === false) {
            settings.url = appendQuery(settings.url, "_=" + Date.now());
        }
        var dataType = settings.dataType, hasPlaceholder = /=\?/.test(settings.url);
        if (dataType == "jsonp" || hasPlaceholder) {
            if (!hasPlaceholder) {
                settings.url = appendQuery(settings.url, "callback=?");
            }
            return $.ajaxJSONP(settings);
        }
        var mime = settings.accepts[dataType], baseHeaders = {}, protocol = /^([\w-]+:)\/\//.test(settings.url) ? RegExp.$1 : window.location.protocol, xhr = settings.xhr(), abortTimeout;
        if (!settings.crossDomain) {
            baseHeaders["X-Requested-With"] = "XMLHttpRequest";
        }
        if (mime) {
            baseHeaders.Accept = mime;
            if (mime.indexOf(",") > -1) {
                mime = mime.split(",", 2)[0];
            }
            xhr.overrideMimeType && xhr.overrideMimeType(mime);
        }
        if (settings.contentType || (settings.contentType !== false && settings.data && settings.type.toUpperCase() != "GET")) {
            baseHeaders["Content-Type"] = (settings.contentType || "application/x-www-form-urlencoded");
        }
        settings.headers = $.extend(baseHeaders, settings.headers || {});
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4) {
                xhr.onreadystatechange = empty;
                clearTimeout(abortTimeout);
                var result, error = false;
                if ((xhr.status >= 200 && xhr.status < 300) || xhr.status == 304 || (xhr.status == 0 && protocol == "file:")) {
                    dataType = dataType || mimeToDataType(xhr.getResponseHeader("content-type"));
                    result = xhr.responseText;
                    try {
                        if (dataType == "script") {
                            (1, eval)(result);
                        } else {
                            if (dataType == "xml") {
                                result = xhr.responseXML;
                            } else {
                                if (dataType == "json") {
                                    result = blankRE.test(result) ? null : $.parseJSON(result);
                                }
                            }
                        }
                    } catch (e) {
                        error = e;
                    }
                    if (error) {
                        ajaxError(error, "parsererror", xhr, settings);
                    } else {
                        ajaxSuccess(result, xhr, settings);
                    }
                } else {
                    ajaxError(null, xhr.status ? "error" : "abort", xhr, settings);
                }
            }
        };
        var async = "async" in settings ? settings.async : true;
        xhr.open(settings.type, settings.url, async);
        for (name in settings.headers) {
            xhr.setRequestHeader(name, settings.headers[name]);
        }
        if (ajaxBeforeSend(xhr, settings) === false) {
            xhr.abort();
            return false;
        }
        if (settings.timeout > 0) {
            abortTimeout = setTimeout(function() {
                xhr.onreadystatechange = empty;
                xhr.abort();
                ajaxError(null, "timeout", xhr, settings);
            }, settings.timeout);
        }
        xhr.send(settings.data ? settings.data : null);
        return xhr;
    };
    function parseArguments(url, data, success, dataType) {
        var hasData = !$.isFunction(data);
        return {url: url,data: hasData ? data : undefined,success: !hasData ? data : $.isFunction(success) ? success : undefined,dataType: hasData ? dataType || success : success};
    }
    $.get = function(url, data, success, dataType) {
        return $.ajax(parseArguments.apply(null, arguments));
    };
    $.post = function(url, data, success, dataType) {
        var options = parseArguments.apply(null, arguments);
        options.type = "POST";
        return $.ajax(options);
    };
    $.getJSON = function(url, data, success) {
        var options = parseArguments.apply(null, arguments);
        options.dataType = "json";
        return $.ajax(options);
    };
    $.fn.load = function(url, data, success) {
        if (!this.length) {
            return this;
        }
        var self = this, parts = url.split(/\s/), selector, options = parseArguments(url, data, success), callback = options.success;
        if (parts.length > 1) {
            options.url = parts[0], selector = parts[1];
        }
        options.success = function(response) {
            self.html(selector ? $("<div>").html(response.replace(rscript, "")).find(selector) : response);
            callback && callback.apply(self, arguments);
        };
        $.ajax(options);
        return this;
    };
    var escape = encodeURIComponent;
    function serialize(params, obj, traditional, scope) {
        var type, array = $.isArray(obj);
        $.each(obj, function(key, value) {
            type = $.type(value);
            if (scope) {
                key = traditional ? scope : scope + "[" + (array ? "" : key) + "]";
            }
            if (!scope && array) {
                params.add(value.name, value.value);
            } else {
                if (type == "array" || (!traditional && type == "object")) {
                    serialize(params, value, traditional, key);
                } else {
                    params.add(key, value);
                }
            }
        });
    }
    $.param = function(obj, traditional) {
        var params = [];
        params.add = function(k, v) {
            this.push(escape(k) + "=" + escape(v));
        };
        serialize(params, obj, traditional);
        return params.join("&").replace(/%20/g, "+");
    };
})(Zepto);
(function(h) {
    var o = h.zepto.qsa, b = {}, n = 1, q = {};
    q.click = q.mousedown = q.mouseup = q.mousemove = "MouseEvents";
    function l(r) {
        return r._zid || (r._zid = n++);
    }
    function c(s, u, t, r) {
        u = e(u);
        if (u.ns) {
            var v = k(u.ns);
        }
        return (b[l(s)] || []).filter(function(w) {
            return w && (!u.e || w.e == u.e) && (!u.ns || v.test(w.ns)) && (!t || l(w.fn) === l(t)) && (!r || w.sel == r);
        });
    }
    function e(r) {
        var s = ("" + r).split(".");
        return {e: s[0],ns: s.slice(1).sort().join(" ")};
    }
    function k(r) {
        return new RegExp("(?:^| )" + r.replace(" ", " .* ?") + "(?: |$)");
    }
    function m(r, t, s) {
        if (h.isObject(r)) {
            h.each(r, s);
        } else {
            r.split(/\s/).forEach(function(u) {
                s(u, t);
            });
        }
    }
    function p(v, u, w, s, r, t) {
        t = !!t;
        var y = l(v), x = (b[y] || (b[y] = []));
        m(u, w, function(C, B) {
            var A = r && r(B, C), E = A || B;
            var D = function(G) {
                var F = E.apply(v, [G].concat(G.data));
                if (F === false) {
                    G.preventDefault();
                }
                return F;
            };
            var z = h.extend(e(C), {fn: B,proxy: D,sel: s,del: A,i: x.length});
            x.push(z);
            v.addEventListener(z.e, D, t);
        });
    }
    function g(t, s, u, r) {
        var v = l(t);
        m(s || "", u, function(x, w) {
            c(t, x, w, r).forEach(function(y) {
                delete b[v][y.i];
                t.removeEventListener(y.e, y.proxy, false);
            });
        });
    }
    h.event = {add: p,remove: g};
    h.proxy = function(t, s) {
        if (h.isFunction(t)) {
            var r = function() {
                return t.apply(s, arguments);
            };
            r._zid = l(t);
            return r;
        } else {
            if (typeof s == "string") {
                return h.proxy(t[s], t);
            } else {
                throw new TypeError("expected function");
            }
        }
    };
    h.fn.bind = function(r, s) {
        return this.each(function() {
            p(this, r, s);
        });
    };
    h.fn.unbind = function(r, s) {
        return this.each(function() {
            g(this, r, s);
        });
    };
    h.fn.one = function(r, s) {
        return this.each(function(u, t) {
            p(this, r, s, null, function(w, v) {
                return function() {
                    var x = w.apply(t, arguments);
                    g(t, v, w);
                    return x;
                };
            });
        });
    };
    var d = function() {
        return true;
    }, a = function() {
        return false;
    }, j = {preventDefault: "isDefaultPrevented",stopImmediatePropagation: "isImmediatePropagationStopped",stopPropagation: "isPropagationStopped"};
    function i(s) {
        var r = h.extend({originalEvent: s}, s);
        h.each(j, function(u, t) {
            r[u] = function() {
                this[t] = d;
                return s[u].apply(s, arguments);
            };
            r[t] = a;
        });
        return r;
    }
    function f(s) {
        if (!("defaultPrevented" in s)) {
            s.defaultPrevented = false;
            var r = s.preventDefault;
            s.preventDefault = function() {
                this.defaultPrevented = true;
                r.call(this);
            };
        }
    }
    h.fn.delegate = function(r, t, u) {
        var s = false;
        if (t == "blur" || t == "focus") {
            if (h.iswebkit) {
                t = t == "blur" ? "focusout" : t == "focus" ? "focusin" : t;
            } else {
                s = true;
            }
        }
        return this.each(function(w, v) {
            p(v, t, u, r, function(x) {
                return function(A) {
                    var y, z = h(A.target).closest(r, v).get(0);
                    if (z) {
                        y = h.extend(i(A), {currentTarget: z,liveFired: v});
                        return x.apply(z, [y].concat([].slice.call(arguments, 1)));
                    }
                };
            }, s);
        });
    };
    h.fn.undelegate = function(r, s, t) {
        return this.each(function() {
            g(this, s, t, r);
        });
    };
    h.fn.live = function(r, s) {
        h(document.body).delegate(this.selector, r, s);
        return this;
    };
    h.fn.die = function(r, s) {
        h(document.body).undelegate(this.selector, r, s);
        return this;
    };
    h.fn.on = function(s, r, t) {
        return r == undefined || h.isFunction(r) ? this.bind(s, r) : this.delegate(r, s, t);
    };
    h.fn.off = function(s, r, t) {
        return r == undefined || h.isFunction(r) ? this.unbind(s, r) : this.undelegate(r, s, t);
    };
    h.fn.trigger = function(r, s) {
        if (typeof r == "string") {
            r = h.Event(r);
        }
        f(r);
        r.data = s;
        return this.each(function() {
            if ("dispatchEvent" in this) {
                this.dispatchEvent(r);
            }
        });
    };
    h.fn.triggerHandler = function(s, t) {
        var u, r;
        this.each(function(w, v) {
            u = i(typeof s == "string" ? h.Event(s) : s);
            u.data = t;
            u.target = v;
            h.each(c(v, s.type || s), function(x, y) {
                r = y.proxy(u);
                if (u.isImmediatePropagationStopped()) {
                    return false;
                }
            });
        });
        return r;
    };
    ("focusin focusout load resize scroll unload click dblclick mousedown mouseup mousemove mouseover mouseout change select keydown keypress keyup error").split(" ").forEach(function(r) {
        h.fn[r] = function(s) {
            return this.bind(r, s);
        };
    });
    ["focus", "blur"].forEach(function(r) {
        h.fn[r] = function(t) {
            if (t) {
                this.bind(r, t);
            } else {
                if (this.length) {
                    try {
                        this.get(0)[r]();
                    } catch (s) {
                    }
                }
            }
            return this;
        };
    });
    h.Event = function(u, t) {
        var v = document.createEvent(q[u] || "Events"), r = true;
        if (t) {
            for (var s in t) {
                (s == "bubbles") ? (r = !!t[s]) : (v[s] = t[s]);
            }
        }
        v.initEvent(u, r, true, null, null, null, null, null, null, null, null, null, null, null, null);
        return v;
    };
})(Zepto);
(function(b) {
    function a(c) {
        var f = this.os = {}, g = this.browser = {}, l = c.match(/WebKit\/([\d.]+)/), e = c.match(/(Android).*?([\d.]+)/) || /HTC/.test(c), m = c.match(/(iPad).*OS\s([\d_]+)/), k = !m && c.match(/(iPhone\sOS)\s([\d_]+)/), n = c.match(/(webOS|hpwOS)[\s\/]([\d.]+)/), j = n && c.match(/TouchPad/), i = c.match(/Kindle\/([\d.]+)/), h = c.match(/Silk\/([\d._]+)/), d = c.match(/(BlackBerry).*Version\/([\d.]+)/);
        if (g.webkit = !!l) {
            g.version = l[1];
        }
        if (e) {
            f.android = true, f.version = e[2];
        }
        if (k) {
            f.ios = f.iphone = true, f.version = k[2].replace(/_/g, ".");
        }
        if (m) {
            f.ios = f.ipad = true, f.version = m[2].replace(/_/g, ".");
        }
        if (n) {
            f.webos = true, f.version = n[2];
        }
        if (j) {
            f.touchpad = true;
        }
        if (d) {
            f.blackberry = true, f.version = d[2];
        }
        if (i) {
            f.kindle = true, f.version = i[1];
        }
        if (h) {
            g.silk = true, g.version = h[1];
        }
        if (!h && f.android && c.match(/Kindle Fire/)) {
            g.silk = true;
        }
    }
    a.call(b, navigator.userAgent);
    b.__detect = a;
})(Zepto);
(function(e, c) {
    var g = "", k, b, i, m = {"": "",Webkit: "webkit",Moz: "",O: "o",ms: "MS"}, j = window.document, d = j.createElement("div"), l = /^((translate|rotate|scale)(X|Y|Z|3d)?|matrix(3d)?|perspective|skew(X|Y)?)$/i, h = {};
    function a(n) {
        return n.toLowerCase();
    }
    function f(n) {
        return k ? k + n : a(n);
    }
    e.each(m, function(o, n) {
        if (d.style[o + "TransitionProperty"] !== c) {
            g = "-" + a(o) + "-";
            k = n;
            return false;
        }
    });
    e.prefix = g;
    h[g + "transition-property"] = h[g + "transition-duration"] = h[g + "transition-timing-function"] = h[g + "animation-name"] = h[g + "animation-duration"] = "";
    e.fx = {off: (k === c && d.style.transitionProperty === c),cssPrefix: g,transitionEnd: f("TransitionEnd"),animationEnd: f("AnimationEnd")};
    e.fn.animate = function(n, o, p, q) {
        if (e.isObject(o)) {
            p = o.easing, q = o.complete, o = o.duration;
        }
        if (o) {
            o = o / 1000;
        }
        return this.anim(n, o, p, q);
    };
    e.fn.anim = function(s, p, o, u) {
        var r, w = {}, t, q = this, n, v = e.fx.transitionEnd;
        if (p === c) {
            p = 0.4;
        }
        if (e.fx.off) {
            p = 0;
        }
        if (typeof s == "string") {
            w[g + "animation-name"] = s;
            w[g + "animation-duration"] = p + "s";
            v = e.fx.animationEnd;
        } else {
            for (t in s) {
                if (l.test(t)) {
                    r || (r = []);
                    r.push(t + "(" + s[t] + ")");
                    delete s[t];
                } else {
                    w[t] = s[t];
                }
            }
            if (r) {
                w[g + "transform"] = r.join(" ");
                s.transform = "transform";
            }
            if (!e.fx.off && typeof s === "object") {
                w[g + "transition-property"] = Object.keys(s).join(", ");
                w[g + "transition-duration"] = p + "s";
                w[g + "transition-timing-function"] = (o || "linear");
            }
        }
        n = function(x) {
            if (typeof x !== "undefined") {
                if (x.target !== x.currentTarget) {
                    return;
                }
                e(x.target).unbind(v, arguments.callee);
            }
            e(this).css(h);
            u && u.call(this);
        };
        if (p > 0) {
            this.bind(v, n);
        }
        setTimeout(function() {
            q.css(w);
            if (p <= 0) {
                setTimeout(function() {
                    q.each(function() {
                        n.call(this);
                    });
                }, 0);
            }
        }, 0);
        return this;
    };
    d = null;
})(Zepto);
(function() {
    var c = function() {
        return function() {
            return ($.os.ios || ($.os.android && $.os.version >= "4.0")) && (!/UCBrowser|baidubrowser/i.test(navigator.userAgent));
        }();
    };
    var e = "", h, i = {"": "",Webkit: "webkit",Moz: "moz",O: "o",ms: "ms"}, g = window.document, b = g.createElement("div");
    function a(j) {
        return j.toLowerCase();
    }
    function f(j) {
        return h ? h + j : a(j);
    }
    $.each(i, function(k, j) {
        if (b.style[k + "TransitionProperty"] !== undefined) {
            e = "-" + a(k) + "-";
            h = j;
            return false;
        }
    });
    function d(m, j) {
        var l = null;
        function k() {
            clearTimeout(l);
            m && m();
            m = null;
        }
        l = setTimeout(k, j || 300);
        return k;
    }
    $.extend($.fn, {page: function(k) {
            var n = this, l = function() {
            }, q = k.init || l, m = k.active || l, p = k.unactive || l, r = k.cancel || l, j = k.beforeUnActive || l, o = false;
            return {thiz: $(n),init: function() {
                    !o && q.apply(this, arguments);
                    o = true;
                },active: function() {
                    m.apply(this, arguments);
                },unactive: p,beforeUnActive: j,cancel: r};
        },pages: function(l) {
            var o = this, p, n = "", q = false, l = $.extend({main: "",side: [],animate: true}, l), k = $("#" + l.main), j = function(r) {
                window.scroll(0, 1);
                if (q) {
                    return;
                }
                q = true;
                k.show();
                r.thiz.animate({translate3d: p + "px,0,0"}, {duration: 250,easing: "ease-out",complete: d(function() {
                        r.thiz.removeClass("active").hide();
                        q = false;
                    })});
                r.beforeUnActive();
                n = "";
            }, m = function(s, r) {
                if (q) {
                    return;
                }
                q = true;
                n = s;
                $.each(l.side, function(t, u, v) {
                    var x = u.thiz;
                    var w = null;
                    var y = function() {
                        k.hide();
                        q = false;
                    };
                    if (s.indexOf(x.attr("id")) == 0) {
                        x.addClass("active").show();
                        u.active();
                        window.scrollTo(0, 1);
                        x.css("-" + h + "-transform", "translate3d(" + p + "px,0,0)").animate({translate3d: "0,0,0"}, {duration: 250,easing: "ease-out",complete: d(y)});
                        return false;
                    }
                });
                !r && l.supportPushState && window.history.pushState({page: s}, s);
            };
            p = $(window).width();
            l.supportPushState = !!window.history.pushState && c();
            $.each(l.side, function(r, s) {
                s.thiz.show();
                s.init();
                s.thiz.hide();
                var t = s.unactive;
                s.unactive = function() {
                    if (t.apply(s, arguments) !== false) {
                        if (l.supportPushState) {
                            window.history.back();
                            s.beforeUnActive();
                        } else {
                            j(s);
                        }
                    }
                };
                s.thiz.find(".back").on("click", function(u) {
                    u.preventDefault();
                    if (s.cancel() === false) {
                        return;
                    }
                    if (l.supportPushState) {
                        window.history.back();
                        s.beforeUnActive();
                    } else {
                        j(s);
                    }
                });
            });
            $(window).on("popstate", function(r) {
                if (r.state && r.state.page) {
                    m(r.state.page, true);
                    return;
                }
                if (!n || $.hash() != "#") {
                    return;
                }
                $.each(l.side, function(s, t) {
                    t.thiz.attr("id") == n && j(t);
                });
            }).on("load", function(r) {
                window.history.replaceState({page: ""}, "");
            });
            return {toPage: m};
        }});
})();
(function() {
    var w = this;
    var k = w._;
    var D = {};
    var C = Array.prototype, f = Object.prototype, r = Function.prototype;
    var H = C.push, o = C.slice, y = C.concat, d = f.toString, j = f.hasOwnProperty;
    var L = C.forEach, q = C.map, E = C.reduce, c = C.reduceRight, b = C.filter, B = C.every, p = C.some, n = C.indexOf, l = C.lastIndexOf, u = Array.isArray, e = Object.keys, F = r.bind;
    var M = function(N) {
        if (N instanceof M) {
            return N;
        }
        if (!(this instanceof M)) {
            return new M(N);
        }
        this._wrapped = N;
    };
    if (typeof exports !== "undefined") {
        if (typeof module !== "undefined" && module.exports) {
            exports = module.exports = M;
        }
        exports._ = M;
    } else {
        w._ = M;
    }
    M.VERSION = "1.4.3";
    var I = M.each = M.forEach = function(S, R, Q) {
        if (S == null) {
            return;
        }
        if (L && S.forEach === L) {
            S.forEach(R, Q);
        } else {
            if (S.length === +S.length) {
                for (var P = 0, N = S.length; P < N; P++) {
                    if (R.call(Q, S[P], P, S) === D) {
                        return;
                    }
                }
            } else {
                for (var O in S) {
                    if (M.has(S, O)) {
                        if (R.call(Q, S[O], O, S) === D) {
                            return;
                        }
                    }
                }
            }
        }
    };
    M.map = M.collect = function(Q, P, O) {
        var N = [];
        if (Q == null) {
            return N;
        }
        if (q && Q.map === q) {
            return Q.map(P, O);
        }
        I(Q, function(T, R, S) {
            N[N.length] = P.call(O, T, R, S);
        });
        return N;
    };
    var g = "Reduce of empty array with no initial value";
    M.reduce = M.foldl = M.inject = function(R, Q, N, P) {
        var O = arguments.length > 2;
        if (R == null) {
            R = [];
        }
        if (E && R.reduce === E) {
            if (P) {
                Q = M.bind(Q, P);
            }
            return O ? R.reduce(Q, N) : R.reduce(Q);
        }
        I(R, function(U, S, T) {
            if (!O) {
                N = U;
                O = true;
            } else {
                N = Q.call(P, N, U, S, T);
            }
        });
        if (!O) {
            throw new TypeError(g);
        }
        return N;
    };
    M.reduceRight = M.foldr = function(T, Q, N, P) {
        var O = arguments.length > 2;
        if (T == null) {
            T = [];
        }
        if (c && T.reduceRight === c) {
            if (P) {
                Q = M.bind(Q, P);
            }
            return O ? T.reduceRight(Q, N) : T.reduceRight(Q);
        }
        var S = T.length;
        if (S !== +S) {
            var R = M.keys(T);
            S = R.length;
        }
        I(T, function(W, U, V) {
            U = R ? R[--S] : --S;
            if (!O) {
                N = T[U];
                O = true;
            } else {
                N = Q.call(P, N, T[U], U, V);
            }
        });
        if (!O) {
            throw new TypeError(g);
        }
        return N;
    };
    M.find = M.detect = function(Q, P, O) {
        var N;
        A(Q, function(T, R, S) {
            if (P.call(O, T, R, S)) {
                N = T;
                return true;
            }
        });
        return N;
    };
    M.filter = M.select = function(Q, P, O) {
        var N = [];
        if (Q == null) {
            return N;
        }
        if (b && Q.filter === b) {
            return Q.filter(P, O);
        }
        I(Q, function(T, R, S) {
            if (P.call(O, T, R, S)) {
                N[N.length] = T;
            }
        });
        return N;
    };
    M.reject = function(P, O, N) {
        return M.filter(P, function(S, Q, R) {
            return !O.call(N, S, Q, R);
        }, N);
    };
    M.every = M.all = function(Q, P, O) {
        P || (P = M.identity);
        var N = true;
        if (Q == null) {
            return N;
        }
        if (B && Q.every === B) {
            return Q.every(P, O);
        }
        I(Q, function(T, R, S) {
            if (!(N = N && P.call(O, T, R, S))) {
                return D;
            }
        });
        return !!N;
    };
    var A = M.some = M.any = function(Q, P, O) {
        P || (P = M.identity);
        var N = false;
        if (Q == null) {
            return N;
        }
        if (p && Q.some === p) {
            return Q.some(P, O);
        }
        I(Q, function(T, R, S) {
            if (N || (N = P.call(O, T, R, S))) {
                return D;
            }
        });
        return !!N;
    };
    M.contains = M.include = function(O, N) {
        if (O == null) {
            return false;
        }
        if (n && O.indexOf === n) {
            return O.indexOf(N) != -1;
        }
        return A(O, function(P) {
            return P === N;
        });
    };
    M.invoke = function(O, P) {
        var N = o.call(arguments, 2);
        return M.map(O, function(Q) {
            return (M.isFunction(P) ? P : Q[P]).apply(Q, N);
        });
    };
    M.pluck = function(O, N) {
        return M.map(O, function(P) {
            return P[N];
        });
    };
    M.where = function(O, N) {
        if (M.isEmpty(N)) {
            return [];
        }
        return M.filter(O, function(Q) {
            for (var P in N) {
                if (N[P] !== Q[P]) {
                    return false;
                }
            }
            return true;
        });
    };
    M.max = function(Q, P, O) {
        if (!P && M.isArray(Q) && Q[0] === +Q[0] && Q.length < 65535) {
            return Math.max.apply(Math, Q);
        }
        if (!P && M.isEmpty(Q)) {
            return -Infinity;
        }
        var N = {computed: -Infinity,value: -Infinity};
        I(Q, function(U, R, T) {
            var S = P ? P.call(O, U, R, T) : U;
            S >= N.computed && (N = {value: U,computed: S});
        });
        return N.value;
    };
    M.min = function(Q, P, O) {
        if (!P && M.isArray(Q) && Q[0] === +Q[0] && Q.length < 65535) {
            return Math.min.apply(Math, Q);
        }
        if (!P && M.isEmpty(Q)) {
            return Infinity;
        }
        var N = {computed: Infinity,value: Infinity};
        I(Q, function(U, R, T) {
            var S = P ? P.call(O, U, R, T) : U;
            S < N.computed && (N = {value: U,computed: S});
        });
        return N.value;
    };
    M.shuffle = function(Q) {
        var P;
        var O = 0;
        var N = [];
        I(Q, function(R) {
            P = M.random(O++);
            N[O - 1] = N[P];
            N[P] = R;
        });
        return N;
    };
    var a = function(N) {
        return M.isFunction(N) ? N : function(O) {
            return O[N];
        };
    };
    M.sortBy = function(Q, P, N) {
        var O = a(P);
        return M.pluck(M.map(Q, function(T, R, S) {
            return {value: T,index: R,criteria: O.call(N, T, R, S)};
        }).sort(function(U, T) {
            var S = U.criteria;
            var R = T.criteria;
            if (S !== R) {
                if (S > R || S === void 0) {
                    return 1;
                }
                if (S < R || R === void 0) {
                    return -1;
                }
            }
            return U.index < T.index ? -1 : 1;
        }), "value");
    };
    var t = function(S, R, O, Q) {
        var N = {};
        var P = a(R || M.identity);
        I(S, function(V, T) {
            var U = P.call(O, V, T, S);
            Q(N, U, V);
        });
        return N;
    };
    M.groupBy = function(P, O, N) {
        return t(P, O, N, function(Q, R, S) {
            (M.has(Q, R) ? Q[R] : (Q[R] = [])).push(S);
        });
    };
    M.countBy = function(P, O, N) {
        return t(P, O, N, function(Q, R) {
            if (!M.has(Q, R)) {
                Q[R] = 0;
            }
            Q[R]++;
        });
    };
    M.sortedIndex = function(U, T, Q, P) {
        Q = Q == null ? M.identity : a(Q);
        var S = Q.call(P, T);
        var N = 0, R = U.length;
        while (N < R) {
            var O = (N + R) >>> 1;
            Q.call(P, U[O]) < S ? N = O + 1 : R = O;
        }
        return N;
    };
    M.toArray = function(N) {
        if (!N) {
            return [];
        }
        if (M.isArray(N)) {
            return o.call(N);
        }
        if (N.length === +N.length) {
            return M.map(N, M.identity);
        }
        return M.values(N);
    };
    M.size = function(N) {
        if (N == null) {
            return 0;
        }
        return (N.length === +N.length) ? N.length : M.keys(N).length;
    };
    M.first = M.head = M.take = function(P, O, N) {
        if (P == null) {
            return void 0;
        }
        return (O != null) && !N ? o.call(P, 0, O) : P[0];
    };
    M.initial = function(P, O, N) {
        return o.call(P, 0, P.length - ((O == null) || N ? 1 : O));
    };
    M.last = function(P, O, N) {
        if (P == null) {
            return void 0;
        }
        if ((O != null) && !N) {
            return o.call(P, Math.max(P.length - O, 0));
        } else {
            return P[P.length - 1];
        }
    };
    M.rest = M.tail = M.drop = function(P, O, N) {
        return o.call(P, (O == null) || N ? 1 : O);
    };
    M.compact = function(N) {
        return M.filter(N, M.identity);
    };
    var x = function(O, P, N) {
        I(O, function(Q) {
            if (M.isArray(Q)) {
                P ? H.apply(N, Q) : x(Q, P, N);
            } else {
                N.push(Q);
            }
        });
        return N;
    };
    M.flatten = function(O, N) {
        return x(O, N, []);
    };
    M.without = function(N) {
        return M.difference(N, o.call(arguments, 1));
    };
    M.uniq = M.unique = function(T, S, R, Q) {
        if (M.isFunction(S)) {
            Q = R;
            R = S;
            S = false;
        }
        var O = R ? M.map(T, R, Q) : T;
        var P = [];
        var N = [];
        I(O, function(V, U) {
            if (S ? (!U || N[N.length - 1] !== V) : !M.contains(N, V)) {
                N.push(V);
                P.push(T[U]);
            }
        });
        return P;
    };
    M.union = function() {
        return M.uniq(y.apply(C, arguments));
    };
    M.intersection = function(O) {
        var N = o.call(arguments, 1);
        return M.filter(M.uniq(O), function(P) {
            return M.every(N, function(Q) {
                return M.indexOf(Q, P) >= 0;
            });
        });
    };
    M.difference = function(O) {
        var N = y.apply(C, o.call(arguments, 1));
        return M.filter(O, function(P) {
            return !M.contains(N, P);
        });
    };
    M.zip = function() {
        var N = o.call(arguments);
        var Q = M.max(M.pluck(N, "length"));
        var P = new Array(Q);
        for (var O = 0; O < Q; O++) {
            P[O] = M.pluck(N, "" + O);
        }
        return P;
    };
    M.object = function(R, P) {
        if (R == null) {
            return {};
        }
        var N = {};
        for (var Q = 0, O = R.length; Q < O; Q++) {
            if (P) {
                N[R[Q]] = P[Q];
            } else {
                N[R[Q][0]] = R[Q][1];
            }
        }
        return N;
    };
    M.indexOf = function(R, P, Q) {
        if (R == null) {
            return -1;
        }
        var O = 0, N = R.length;
        if (Q) {
            if (typeof Q == "number") {
                O = (Q < 0 ? Math.max(0, N + Q) : Q);
            } else {
                O = M.sortedIndex(R, P);
                return R[O] === P ? O : -1;
            }
        }
        if (n && R.indexOf === n) {
            return R.indexOf(P, Q);
        }
        for (; O < N; O++) {
            if (R[O] === P) {
                return O;
            }
        }
        return -1;
    };
    M.lastIndexOf = function(R, P, Q) {
        if (R == null) {
            return -1;
        }
        var N = Q != null;
        if (l && R.lastIndexOf === l) {
            return N ? R.lastIndexOf(P, Q) : R.lastIndexOf(P);
        }
        var O = (N ? Q : R.length);
        while (O--) {
            if (R[O] === P) {
                return O;
            }
        }
        return -1;
    };
    M.range = function(S, Q, R) {
        if (arguments.length <= 1) {
            Q = S || 0;
            S = 0;
        }
        R = arguments[2] || 1;
        var O = Math.max(Math.ceil((Q - S) / R), 0);
        var N = 0;
        var P = new Array(O);
        while (N < O) {
            P[N++] = S;
            S += R;
        }
        return P;
    };
    var G = function() {
    };
    M.bind = function(Q, O) {
        var N, P;
        if (Q.bind === F && F) {
            return F.apply(Q, o.call(arguments, 1));
        }
        if (!M.isFunction(Q)) {
            throw new TypeError;
        }
        N = o.call(arguments, 2);
        return P = function() {
            if (!(this instanceof P)) {
                return Q.apply(O, N.concat(o.call(arguments)));
            }
            G.prototype = Q.prototype;
            var S = new G;
            G.prototype = null;
            var R = Q.apply(S, N.concat(o.call(arguments)));
            if (Object(R) === R) {
                return R;
            }
            return S;
        };
    };
    M.bindAll = function(O) {
        var N = o.call(arguments, 1);
        if (N.length === 0) {
            N = M.functions(O);
        }
        I(N, function(P) {
            O[P] = M.bind(O[P], O);
        });
        return O;
    };
    M.memoize = function(P, O) {
        var N = {};
        O || (O = M.identity);
        return function() {
            var Q = O.apply(this, arguments);
            return M.has(N, Q) ? N[Q] : (N[Q] = P.apply(this, arguments));
        };
    };
    M.delay = function(O, P) {
        var N = o.call(arguments, 2);
        return setTimeout(function() {
            return O.apply(null, N);
        }, P);
    };
    M.defer = function(N) {
        return M.delay.apply(M, [N, 1].concat(o.call(arguments, 1)));
    };
    M.throttle = function(S, U) {
        var Q, P, T, N;
        var R = 0;
        var O = function() {
            R = new Date;
            T = null;
            N = S.apply(Q, P);
        };
        return function() {
            var V = new Date;
            var W = U - (V - R);
            Q = this;
            P = arguments;
            if (W <= 0) {
                clearTimeout(T);
                T = null;
                R = V;
                N = S.apply(Q, P);
            } else {
                if (!T) {
                    T = setTimeout(O, W);
                }
            }
            return N;
        };
    };
    M.debounce = function(P, R, O) {
        var Q, N;
        return function() {
            var V = this, U = arguments;
            var T = function() {
                Q = null;
                if (!O) {
                    N = P.apply(V, U);
                }
            };
            var S = O && !Q;
            clearTimeout(Q);
            Q = setTimeout(T, R);
            if (S) {
                N = P.apply(V, U);
            }
            return N;
        };
    };
    M.once = function(P) {
        var N = false, O;
        return function() {
            if (N) {
                return O;
            }
            N = true;
            O = P.apply(this, arguments);
            P = null;
            return O;
        };
    };
    M.wrap = function(N, O) {
        return function() {
            var P = [N];
            H.apply(P, arguments);
            return O.apply(this, P);
        };
    };
    M.compose = function() {
        var N = arguments;
        return function() {
            var O = arguments;
            for (var P = N.length - 1; P >= 0; P--) {
                O = [N[P].apply(this, O)];
            }
            return O[0];
        };
    };
    M.after = function(O, N) {
        if (O <= 0) {
            return N();
        }
        return function() {
            if (--O < 1) {
                return N.apply(this, arguments);
            }
        };
    };
    M.keys = e || function(P) {
        if (P !== Object(P)) {
            throw new TypeError("Invalid object");
        }
        var O = [];
        for (var N in P) {
            if (M.has(P, N)) {
                O[O.length] = N;
            }
        }
        return O;
    };
    M.values = function(P) {
        var N = [];
        for (var O in P) {
            if (M.has(P, O)) {
                N.push(P[O]);
            }
        }
        return N;
    };
    M.pairs = function(P) {
        var O = [];
        for (var N in P) {
            if (M.has(P, N)) {
                O.push([N, P[N]]);
            }
        }
        return O;
    };
    M.invert = function(P) {
        var N = {};
        for (var O in P) {
            if (M.has(P, O)) {
                N[P[O]] = O;
            }
        }
        return N;
    };
    M.functions = M.methods = function(P) {
        var O = [];
        for (var N in P) {
            if (M.isFunction(P[N])) {
                O.push(N);
            }
        }
        return O.sort();
    };
    M.extend = function(N) {
        I(o.call(arguments, 1), function(O) {
            if (O) {
                for (var P in O) {
                    N[P] = O[P];
                }
            }
        });
        return N;
    };
    M.pick = function(O) {
        var P = {};
        var N = y.apply(C, o.call(arguments, 1));
        I(N, function(Q) {
            if (Q in O) {
                P[Q] = O[Q];
            }
        });
        return P;
    };
    M.omit = function(P) {
        var Q = {};
        var O = y.apply(C, o.call(arguments, 1));
        for (var N in P) {
            if (!M.contains(O, N)) {
                Q[N] = P[N];
            }
        }
        return Q;
    };
    M.defaults = function(N) {
        I(o.call(arguments, 1), function(O) {
            if (O) {
                for (var P in O) {
                    if (N[P] == null) {
                        N[P] = O[P];
                    }
                }
            }
        });
        return N;
    };
    M.clone = function(N) {
        if (!M.isObject(N)) {
            return N;
        }
        return M.isArray(N) ? N.slice() : M.extend({}, N);
    };
    M.tap = function(O, N) {
        N(O);
        return O;
    };
    var J = function(U, T, O, P) {
        if (U === T) {
            return U !== 0 || 1 / U == 1 / T;
        }
        if (U == null || T == null) {
            return U === T;
        }
        if (U instanceof M) {
            U = U._wrapped;
        }
        if (T instanceof M) {
            T = T._wrapped;
        }
        var R = d.call(U);
        if (R != d.call(T)) {
            return false;
        }
        switch (R) {
            case "[object String]":
                return U == String(T);
            case "[object Number]":
                return U != +U ? T != +T : (U == 0 ? 1 / U == 1 / T : U == +T);
            case "[object Date]":
            case "[object Boolean]":
                return +U == +T;
            case "[object RegExp]":
                return U.source == T.source && U.global == T.global && U.multiline == T.multiline && U.ignoreCase == T.ignoreCase;
        }
        if (typeof U != "object" || typeof T != "object") {
            return false;
        }
        var N = O.length;
        while (N--) {
            if (O[N] == U) {
                return P[N] == T;
            }
        }
        O.push(U);
        P.push(T);
        var W = 0, X = true;
        if (R == "[object Array]") {
            W = U.length;
            X = W == T.length;
            if (X) {
                while (W--) {
                    if (!(X = J(U[W], T[W], O, P))) {
                        break;
                    }
                }
            }
        } else {
            var S = U.constructor, Q = T.constructor;
            if (S !== Q && !(M.isFunction(S) && (S instanceof S) && M.isFunction(Q) && (Q instanceof Q))) {
                return false;
            }
            for (var V in U) {
                if (M.has(U, V)) {
                    W++;
                    if (!(X = M.has(T, V) && J(U[V], T[V], O, P))) {
                        break;
                    }
                }
            }
            if (X) {
                for (V in T) {
                    if (M.has(T, V) && !(W--)) {
                        break;
                    }
                }
                X = !W;
            }
        }
        O.pop();
        P.pop();
        return X;
    };
    M.isEqual = function(O, N) {
        return J(O, N, [], []);
    };
    M.isEmpty = function(O) {
        if (O == null) {
            return true;
        }
        if (M.isArray(O) || M.isString(O)) {
            return O.length === 0;
        }
        for (var N in O) {
            if (M.has(O, N)) {
                return false;
            }
        }
        return true;
    };
    M.isElement = function(N) {
        return !!(N && N.nodeType === 1);
    };
    M.isArray = u || function(N) {
        return d.call(N) == "[object Array]";
    };
    M.isObject = function(N) {
        return N === Object(N);
    };
    I(["Arguments", "Function", "String", "Number", "Date", "RegExp"], function(N) {
        M["is" + N] = function(O) {
            return d.call(O) == "[object " + N + "]";
        };
    });
    if (!M.isArguments(arguments)) {
        M.isArguments = function(N) {
            return !!(N && M.has(N, "callee"));
        };
    }
    if (typeof (/./) !== "function") {
        M.isFunction = function(N) {
            return typeof N === "function";
        };
    }
    M.isFinite = function(N) {
        return isFinite(N) && !isNaN(parseFloat(N));
    };
    M.isNaN = function(N) {
        return M.isNumber(N) && N != +N;
    };
    M.isBoolean = function(N) {
        return N === true || N === false || d.call(N) == "[object Boolean]";
    };
    M.isNull = function(N) {
        return N === null;
    };
    M.isUndefined = function(N) {
        return N === void 0;
    };
    M.has = function(O, N) {
        return j.call(O, N);
    };
    M.noConflict = function() {
        w._ = k;
        return this;
    };
    M.identity = function(N) {
        return N;
    };
    M.times = function(R, Q, P) {
        var N = Array(R);
        for (var O = 0; O < R; O++) {
            N[O] = Q.call(P, O);
        }
        return N;
    };
    M.random = function(O, N) {
        if (N == null) {
            N = O;
            O = 0;
        }
        return O + (0 | Math.random() * (N - O + 1));
    };
    var m = {escape: {"&": "&amp;","<": "&lt;",">": "&gt;",'"': "&quot;","'": "&#x27;","/": "&#x2F;"}};
    m.unescape = M.invert(m.escape);
    var K = {escape: new RegExp("[" + M.keys(m.escape).join("") + "]", "g"),unescape: new RegExp("(" + M.keys(m.unescape).join("|") + ")", "g")};
    M.each(["escape", "unescape"], function(N) {
        M[N] = function(O) {
            if (O == null) {
                return "";
            }
            return ("" + O).replace(K[N], function(P) {
                return m[N][P];
            });
        };
    });
    M.result = function(N, P) {
        if (N == null) {
            return null;
        }
        var O = N[P];
        return M.isFunction(O) ? O.call(N) : O;
    };
    M.mixin = function(N) {
        I(M.functions(N), function(O) {
            var P = M[O] = N[O];
            M.prototype[O] = function() {
                var Q = [this._wrapped];
                H.apply(Q, arguments);
                return s.call(this, P.apply(M, Q));
            };
        });
    };
    var z = 0;
    M.uniqueId = function(N) {
        var O = "" + (++z);
        return N ? N + O : O;
    };
    M.templateSettings = {evaluate: /<%([\s\S]+?)%>/g,interpolate: /<%=([\s\S]+?)%>/g,escape: /<%-([\s\S]+?)%>/g};
    var v = /(.)^/;
    var h = {"'": "'","\\": "\\","\r": "r","\n": "n","\t": "t","\u2028": "u2028","\u2029": "u2029"};
    var i = /\\|'|\r|\n|\t|\u2028|\u2029/g;
    M.template = function(V, Q, P) {
        var O;
        P = M.defaults({}, P, M.templateSettings);
        var R = new RegExp([(P.escape || v).source, (P.interpolate || v).source, (P.evaluate || v).source].join("|") + "|$", "g");
        var S = 0;
        var N = "__p+='";
        V.replace(R, function(X, Y, W, aa, Z) {
            N += V.slice(S, Z).replace(i, function(ab) {
                return "\\" + h[ab];
            });
            if (Y) {
                N += "'+\n((__t=(" + Y + "))==null?'':_.escape(__t))+\n'";
            }
            if (W) {
                N += "'+\n((__t=(" + W + "))==null?'':__t)+\n'";
            }
            if (aa) {
                N += "';\n" + aa + "\n__p+='";
            }
            S = Z + X.length;
            return X;
        });
        N += "';\n";
        if (!P.variable) {
            N = "with(obj||{}){\n" + N + "}\n";
        }
        N = "var __t,__p='',__j=Array.prototype.join,print=function(){__p+=__j.call(arguments,'');};\n" + N + "return __p;\n";
        try {
            O = new Function(P.variable || "obj", "_", N);
        } catch (T) {
            T.source = N;
            throw T;
        }
        if (Q) {
            return O(Q, M);
        }
        var U = function(W) {
            return O.call(this, W, M);
        };
        U.source = "function(" + (P.variable || "obj") + "){\n" + N + "}";
        return U;
    };
    M.chain = function(N) {
        return M(N).chain();
    };
    var s = function(N) {
        return this._chain ? M(N).chain() : N;
    };
    M.mixin(M);
    I(["pop", "push", "reverse", "shift", "sort", "splice", "unshift"], function(N) {
        var O = C[N];
        M.prototype[N] = function() {
            var P = this._wrapped;
            O.apply(P, arguments);
            if ((N == "shift" || N == "splice") && P.length === 0) {
                delete P[0];
            }
            return s.call(this, P);
        };
    });
    I(["concat", "join", "slice"], function(N) {
        var O = C[N];
        M.prototype[N] = function() {
            return s.call(this, O.apply(this._wrapped, arguments));
        };
    });
    M.extend(M.prototype, {chain: function() {
            this._chain = true;
            return this;
        },value: function() {
            return this._wrapped;
        }});
}).call(this);
(function() {
    $.extend($, {getDay: function(d) {
            var c = d.split("-"), b = new Date(c[0], c[1] - 1, c[2]);
            b = b.getDay();
            return "星期" + (b == 1 ? "一" : b == 2 ? "二" : b == 3 ? "三" : b == 4 ? "四" : b == 5 ? "五" : b == 6 ? "六" : "日");
        },hash: function(d) {
            var b = (window || this).location.href, c = b.match(/#(.*)$/);
            d && (window.location.href = "#" + d);
            return c ? c[1] : "#";
        },lastHref: function() {
            var b = window.location.href.split("/");
            return b[b.length - 1];
        },jumpDate: function(f, g) {
            var b = f.split("-"), e = new Date(b[0], b[1] - 1, b[2]);
            e.setDate(e.getDate() + g);
            var c = e.getMonth() + 1;
            return e.getFullYear() + "-" + (c < 10 ? "0" + c : c) + "-" + (e.getDate() < 10 ? "0" + e.getDate() : e.getDate());
        },href: function() {
            var d = this, e = window.location.href, c = new RegExp(), b;
            e = e.replace(/#.*/, "");
            return {getValue: function() {
                    return e;
                },replace: function() {
                    e = e.replace(arguments[0], arguments[1]);
                    return this;
                },append: function(f) {
                    e = e + "&" + f;
                    return this;
                },path: function(f) {
                    e = e.indexOf("?") != -1 ? e.replace(/\/[^\/]*\?/, "/" + f + "?") : e.replace(/\/[^/]*$/, "/" + f + "?");
                    return this;
                },param: function() {
                    var g = arguments, f = g.length;
                    if (f == 0) {
                        e = e.replace(/\?(.)*/, "?");
                        return this;
                    }
                    if (f == 1 && typeof g[0] == "string") {
                        c.compile("[&?]" + g[0] + "=([^=&?#]*)");
                        return e.match(c) ? e.match(c)[1] : "";
                    }
                    if (f == 2) {
                        c.compile("[&?]" + g[0] + "=([^=&?#]*)");
                        b = e.match(c);
                        e = b ? e.replace(g[0] + "=" + b[1], g[0] + "=" + g[1]) : (e + (e.indexOf("?") == -1 ? "?" : (e.indexOf("?") == e.length - 1 ? "" : "&")) + g[0] + "=" + g[1]);
                        return this;
                    }
                    if (typeof g[0] == "object") {
                        var h, j = g[0];
                        for (h in j) {
                            g.callee(h, j[h]);
                        }
                    }
                    return this;
                },exec: function() {
                    window.location.href = e;
                }};
        },fixTo8Date: function(f, b, e, c) {
            return [f, b < 10 ? "0" + b : b, e < 10 ? "0" + e : e].join(c || "");
        },getTimeString: function(b) {
            var c = new Date(), e = function(d) {
                return d < 10 ? "0" + d : d;
            };
            c.setTime(b);
            return c.getFullYear() + "-" + e(c.getMonth() + 1) + "-" + e(c.getDate()) + " " + e(c.getHours()) + ":" + e(c.getMinutes());
        },sendClick: function(b) {
            _hmt.push(["_trackEvent", "common", "cnt", b]);
        },throttle: function(b, c, g) {
            var e = 0, d;
            if (typeof c !== "function") {
                g = c;
                c = b;
                b = 250;
            }
            function f() {
                var k = this, l = Date.now() - e, j = arguments;
                function i() {
                    e = Date.now();
                    c.apply(k, j);
                }
                function h() {
                    d = undefined;
                }
                if (g && !d) {
                    i();
                }
                d && clearTimeout(d);
                if (g === undefined && l > b) {
                    i();
                } else {
                    d = setTimeout(g ? h : i, g === undefined ? b - l : b);
                }
            }
            f._zid = c._zid = c._zid || $.proxy(c)._zid;
            return f;
        },debounce: function(b, d, c) {
            return d === undefined ? $.throttle(250, b, false) : $.throttle(b, d, c === undefined ? false : c !== false);
        },getDateString: function(b) {
            var c = b, e = function(d) {
                return d < 10 ? "0" + d : d;
            };
            return c.getFullYear() + "-" + e(c.getMonth() + 1) + "-" + e(c.getDate());
        }});
    $.extend($.fn, {countdown: function(c) {
            var d = this, g = c.seconds || 10, f = c.text || "", b = c.url || window.location.href, e = setInterval(function() {
                if (g <= 1) {
                    clearInterval(e);
                    window.location.href = b;
                    return;
                }
                $(d).html(--g + f);
            }, 1000);
            $(d).html(g + f);
            return this;
        },hover: function(c) {
            var b = $(this), c = c || "hover";
            b.each(function(d, e) {
                $(e).on("touchstart", function() {
                    var f = setTimeout(function() {
                        $(e).addClass(c);
                    }, 50);
                    $(document).one("touchmove touchend", function() {
                        clearTimeout(f);
                        $(e).removeClass(c);
                    });
                });
            });
        },textfold: function(b) {
            var c = $.extend({len: 90,text: "",style: {width: 30,height: 20},clz: "qn_arrow_grey b qn_fr"}, b);
            $.each($(this), function() {
                var f = this;
                var d = $(f).html();
                if (d.length > c.len) {
                    var e = $("<span>").html(c.text).addClass(c.clz).css(c.style);
                    $(f).one("click", function() {
                        $(f).html(d);
                    });
                    $(f).html($(f).html().substring(0, c.len) + "...").append(e);
                }
            });
            return this;
        },simuInput: function() {
            this.each(function(b) {
                var c = $(this);
                c.on("click", function(g) {
                    if (c.hasClass("disable")) {
                        return;
                    }
                    var f = g.target, d;
                    if (f.tagName == "INPUT") {
                        return;
                    }
                    d = $(this).find("input")[0];
                    d.checked = d.checked ? false : true;
                });
            });
            return this;
        },clearInput: function(c) {
            var d = c.clearIcon || "", b = c.delBtnClass || "";
            this.each(function() {
                var f = $(this), e = f.find("input"), g = f.find(d);
                if (e.length == 0) {
                    return;
                }
                e.on("input focus", function(h) {
                    if (this.value == "") {
                        g.removeClass(b);
                    } else {
                        g.addClass(b);
                    }
                }).on("blur", function(h) {
                    setTimeout(function() {
                        g.removeClass(b);
                    }, 200);
                });
                g.on("touchstart", function(h) {
                    h.preventDefault();
                    e.val("");
                    e.blur().focus();
                });
            });
        },switchBtn: function(b) {
            var b = b || {};
            this.each(function(e) {
                var g = $(this), f = g.find(".qn_switch_btn"), d = b.status || 0, c = g.find(".on"), h = g.find(".off");
                if (d == 1) {
                    h.css("visibility", "hidden");
                } else {
                    c.css("visibility", "hidden");
                    f.css("left", "2px");
                    g.addClass(b.className);
                }
                g.on("click", function() {
                    g.toggleClass(b.className);
                    if (g.hasClass(b.className)) {
                        h.css("visibility", "visible");
                        f.animate({left: "2px"}, {duration: 200,easing: "ease-out",complete: function() {
                                g.trigger("off");
                                c.css("visibility", "hidden");
                            }});
                    } else {
                        c.css("visibility", "visible");
                        f.animate({left: "42px"}, {duration: 200,easing: "ease-out",complete: function() {
                                h.css("visibility", "hidden");
                                g.trigger("on");
                            }});
                    }
                });
            });
            return this;
        },fix: function(d) {
            var c = this;
            if (c.attr("isFixed")) {
                return c;
            }
            c.css(d).css("position", "fixed").attr("isFixed", true);
            var g = $('<div style="position:fixed;top:10px;"></div>').appendTo("body"), e = g[0].getBoundingClientRect().top, b = function() {
                if (window.pageYOffset > 0) {
                    if (g[0].getBoundingClientRect().top !== e) {
                        c.css("position", "absolute").attr("isFixed", false);
                        f();
                        $(document).on("scrollStop", f);
                        $(window).on("ortchange", f);
                    }
                    $(document).off("scrollStop", b);
                    g.remove();
                }
            }, f = function() {
                if (c.attr("isClosed")) {
                    return;
                }
                c.show().css({top: window.pageYOffset + (d.bottom !== undefined ? window.innerHeight - c.height() - parseInt(d.bottom, 10) : (d.top || 0)),left: d.right !== undefined ? document.body.offsetWidth - c.width() - d.right : (d.left || 0)});
                d.width == "100%" && c.css("width", document.body.offsetWidth);
            };
            $(document).on("scrollStop", b);
            $(document).on("touchmove", function() {
                if (c.attr("isFixed") == "false") {
                    !c.attr("isClosed") && c.hide();
                }
            });
            return c;
        },paging: function(b) {
            this.each(function(d) {
                var c = this;
                c.opt = $.extend({next: function() {
                    },pre: function() {
                    }}, b);
                var f = $(c).find(".pre");
                var e = $(c).find(".next");
                f.on("click", function(i) {
                    i.preventDefault();
                    var h = $(c).find(".num");
                    var g = h.html();
                    if (g <= 1) {
                        return;
                    }
                    g--;
                    h.html(g);
                    c.opt.pre.call(c);
                });
                e.on("click", function(j) {
                    j.preventDefault();
                    var i = $(c).find(".total");
                    var h = $(c).find(".num");
                    var g = parseInt(h.html(), 10);
                    var k = parseInt(i.html(), 10);
                    if (g >= k) {
                        return;
                    }
                    g++;
                    h.html(g);
                    c.opt.next.call(c);
                });
            });
            return this;
        },fadeIn: function(b) {
            b = b || 350;
            return this.each(function() {
                $(this).css({display: "",opacity: 0}).animate({opacity: 1}, b);
            });
        },fadeOut: function(b) {
            b = b || 350;
            $(this).css({opacity: 1}).animate({opacity: 0}, b).hide();
        }});
    $.has3d = ("WebKitCSSMatrix" in window && "m11" in new WebKitCSSMatrix());
    (function a() {
        $(window).on("scroll", $.debounce(80, function() {
            $(document).trigger("scrollStop");
        }, false));
    })();
})();
(function() {
    var c = 1000 * 60 * 60 * 24;
    var d = function(f) {
        this.opt = $.extend({holiday: [],tiaoxiu: [],start: "",sContent: "",back: "",bContent: "",container: "",validDays: 90,validType: 0,showDay: true,sameDay: true,offset: 2,tags: {},tpl: "#calendarTpl",beforeCallback: function() {
            },chooseCallback: function() {
            }}, f);
        this.today = new Date();
        this._init();
    };
    d.prototype = {constructor: d,_transToDate: function(h) {
            var g = [], f = h;
            if (f instanceof Date) {
                return f;
            }
            if (f instanceof Object) {
                g.push(f.y);
                g.push(f.m);
                g.push(f.d);
            } else {
                if (!arguments[0] || arguments[0] == "null") {
                    return new Date();
                }
                g = arguments.length == 3 ? [arguments[0], arguments[1], arguments[2]] : f.split("-");
            }
            return new Date(g[0], g[1] - 1, g[2]);
        },_transToObj: function(f) {
            var f = this._transToDate(f);
            return {y: f.getFullYear(),m: f.getMonth() + 1,d: f.getDate()};
        },_transToString: function(f) {
            var f = this._transToObj(f);
            return $.fixTo8Date(f.y, f.m, f.d, "-");
        },_compareDay: function(g, f) {
            var g = this._transToDate(g), f = this._transToDate(f);
            return g - f >= 0;
        },_validDay: function() {
            var g = this._transToDate.apply(this, arguments), f = this.today, h = this.opt.validDays - 1;
            return g - f >= -c && g - f <= h * c;
        },_init: function() {
            /*var j = this, h = this.opt, g = this.today, i = new Date(), k = "", f = "";
            i.setDate(g.getDate() + 1);
            this.zStart = $(h.sContent || h.start);
            k = this.zStart.html() || i;
            this.zStart.html(this.getDayString(k));
            $(h.start).on("click", function(l) {
                l.preventDefault();
                j.zActive = j.zStart;
                j.active(j.zStart.html());
                h.beforeCallback.apply(j);
            });
            this.zBack = $(h.bContent || h.back);
            f = this._transToDate(this.zBack.html());
            this.zBack.html(this._compareDay(k, f) ? this.getDayString($.jumpDate(this.zStart.html().split("&")[0], h.offset)) : this.getDayString(f));
            $(h.back).on("click", function() {
                j.zActive = j.zBack;
                j.active(j.zStart.html(), j.zBack.html());
                h.beforeCallback.apply(j);
            });*/
            /*this.zContainer = $(h.container).delegate("td", "touchstart", function() {
                var m = $(this);
                if (m.hasClass("null") || m.hasClass("disable")) {
                    return;
                }
                var l = setTimeout(function() {
                    m.addClass("hover");
                }, 50);
                $(document).one("touchmove touchend", function() {
                    clearTimeout(l);
                    m.removeClass("hover");
                });
            }).delegate("td", "click", function() {
                j.choose(this);
            });
            this.choosenDate = this._transToObj(k);*/
        },render: function(p) {
            /*var l = this, n = this.today, k = function(x, r) {
                week = l._getDay(new Date(x, r - 1, 1));
                week = week == 0 ? 6 : week - 1;
                days = new Date(x, r, 1);
                days.setDate(days.getDate() - 1);
                days = days.getDate();
                var w = [], t = [], v = 0, q = Math.ceil((week + days) / 7), u = 0, s;
                for (; v < q; v++) {
                    t = [];
                    for (u = 0; u < 7; u++) {
                        s = v * 7 + u + 1 - week;
                        v == 0 ? u < week ? t.push(null) : t.push(l._validDay(x, r, s) ? s : -s) : v * 7 + u - week >= days ? t.push(null) : t.push(l._validDay(x, r, s) ? s : -s);
                    }
                    w.push(t);
                }
                return w;
            };
            l.zContainer.html("");
            var o = new Date(), f = l.opt, m = n.getFullYear(), j = n.getMonth();
            o.setDate(n.getDate() + f.validDays);
            for (var g = 0, h = (o.getFullYear() - m) * 12 + o.getMonth() + l.opt.validType - j; g < h; g++) {
                if (j + g > 11) {
                    l.zContainer.append(_.template($(l.opt.tpl).html(), {choosen: this.choosenDate,display: this._transToObj(new Date(m + 1, j + g - 12)),today: this._transToObj(n),weeks: k(m + 1, j + g - 11),holiday: l.opt.holiday,tiaoxiu: l.opt.tiaoxiu,tags: p}));
                } else {
                    l.zContainer.append(_.template($(l.opt.tpl).html(), {choosen: this.choosenDate,display: this._transToObj(new Date(m, j + g)),today: this._transToObj(n),weeks: k(m, j + g + 1),holiday: l.opt.holiday,tiaoxiu: l.opt.tiaoxiu,tags: p}));
                }
            }*/
            return this;
        },fill: function() {
        },choose: function(i) {
            /*if (i.className.match(/disable|null/)) {
                return;
            }
            var h = this, g = h.opt, j = this.opt.chooseCallback, f = $(i).data("day");
            h.zContainer.find(".active").removeClass("active");
            this.choosenDate = h._transToObj(f);
            i.className = "active";
            if (h.zBack.length && h.zActive.attr("id") == h.zStart.attr("id")) {
                h._compareDay(h.choosenDate, h.zBack.html().split("&nbsp;")[0]) && h.zBack.html(h.getDayString($.jumpDate(f, g.offset)));
            }
            h.zActive && h.zActive.html(h.getDayString(h.choosenDate));
            j && j.call(this, i);*/
        },active: function() {
            /*var h = this, i, g = arguments, f = g.length == 1 ? 0 : 1;
            g[0] = g[0].split("&nbsp;")[0];
            if (g.length == 1) {
                i = h.today;
            } else {
                g[1] = g[1].split("&nbsp;")[0];
                i = h._transToDate(g[0]);
                if (!h.opt.sameDay) {
                    i.setDate(i.getDate() + 1);
                }
            }
            /*this.zContainer.find("td").each(function(k, l) {
                var m = l.className, j = $(l);
                if (m.match(/null/) || !j.data("day")) {
                    return true;
                }
                var n = h._transToDate(j.data("day"));
                if (n - i > -c) {
                    j.removeClass("disable");
                } else {
                    j.addClass("disable");
                }
                j.removeClass("active");
                h._transToString(n) == g[f] && j.addClass("active");
            });*/
        },_getDay: function(f) {
            return this._transToDate(f).getDay();
        },_daysOfWeek: function(g) {
            var h = ["日", "一", "二", "三", "四", "五", "六"], f = a(g).getDay();
            return "星期" + h[f];
        },getDayString: function(f) {
            return this.opt.showDay ? this._transToString(f) + "&nbsp;" + this._daysOfWeek(f) : this._transToString(f);
        }};
    $.QunarCalendar = d;
    function a(h) {
        var g = [], f = h;
        if (f instanceof Date) {
            return f;
        }
        if (f instanceof Object) {
            g.push(f.y);
            g.push(f.m);
            g.push(f.d);
        } else {
            if (!arguments[0] || arguments[0] == "null") {
                return new Date();
            }
            g = arguments.length == 3 ? [arguments[0], arguments[1], arguments[2]] : f.split("-");
        }
        return new Date(parseInt(g[0], 10), parseInt(g[1] - 1, 10), parseInt(g[2], 10));
    }
    function e(f) {
        return $.getDateString(f);
    }
    var b = function(f) {
        this.options = $.extend({serverDate: "",container: "",tpl: "#cTpl",validDays: 90,init: function() {
            },fill: function() {
            },active: function() {
            },unactive: function() {
            }}, f);
        this._data = {};
        this._init();
    };
    b.prototype = {constructor: d,_init: function() {
            var g = this, f = g.options;
            this.zContainer = $(f.container).delegate("td", "touchstart", function() {
                var i = $(this);
                if (i.hasClass("null") || i.hasClass("disable")) {
                    return;
                }
                var h = setTimeout(function() {
                    i.addClass("hover");
                }, 50);
                $(document).one("touchmove touchend", function() {
                    clearTimeout(h);
                    i.removeClass("hover");
                });
            }).delegate("td", "click", function() {
                g.unactive(this);
            });
            f.init.call(this);
        },_render: function(g, k) {
            var i = this, f = i.options, h = [], j = f.fill;
            getWeeks = function(u, p, t) {
                var o = new Date(u, p - 1, 1).getDay(), v = new Date(u, p, 1);
                o = o == 0 ? 6 : o - 1;
                v.setDate(v.getDate() - 1);
                v = v.getDate();
                var l = [], w = [], s = Math.ceil((o + v) / 7), r, q, n;
                for (r = 0; r < s; r++) {
                    w = [];
                    for (q = 0; q < 7; q++) {
                        n = r * 7 + q + 1 - o;
                        if (r == 0) {
                            q < o ? w.push({date: e(new Date(u, p - 1, n)),content: {className: "null",content: ""}}) : w.push({date: e(new Date(u, p - 1, n)),content: j.call(i, new Date(u, p - 1, n))});
                        } else {
                            r * 7 + q - o >= v ? w.push({date: e(new Date(u, p - 1, n)),content: {className: "null",content: ""}}) : w.push({date: e(new Date(u, p - 1, n)),content: j.call(i, new Date(u, p - 1, n))});
                        }
                    }
                    l.push(w);
                }
                return l;
            };
            i.zContainer.append(_.template($(f.tpl).html(), {year: g,month: k,weeks: getWeeks(g, k)}));
        },render: function(i, l, k) {
            var j = this, h = j.options, g, f;
            if (k.clear) {
                j.zContainer.html("");
            }
            if (!i && h.validDays) {
                g = a(h.serverDate);
                f = a(h.serverDate);
                f.setDate(g.getDate() + h.validDays);
                while (g - f < 0) {
                    i = g.getFullYear();
                    l = g.getMonth() + 1;
                    j._render(i, l);
                    g.setDate(1);
                    g.setMonth(g.getMonth() + 1);
                }
                return j;
            }
            if (i && l) {
                j._render(i, l);
            }
            return j;
        },active: function(h, g) {
            var i = a(h), f = this.options.active;
            f && f.call(this, i, g);
            return this;
        },unactive: function(h, g) {
            var f = this.options.unactive;
            f && f.call(this, h, g);
            return this;
        },getDay: function(g) {
            var h = ["日", "一", "二", "三", "四", "五", "六"], f = a(g).getDay();
            return "星期" + h[f];
        },attr: function(g, k) {
            var h = this, j = h._data;
            if (typeof g == "string") {
                if (k === undefined) {
                    return j[g];
                }
                j[g] = k;
                return this;
            }
            if (typeof g == "object") {
                for (var f in g) {
                    h.attr(f, g[f]);
                }
                return this;
            }
        },transToDate: a};
    $._QunarCalendar = b;
})(Zepto);

(function() {
    var b = function(f) {
        this.options = $.extend({data: [], day: '', value: '', container: "",tpl: "#typeTpl",init: function() {
            },fill: function() {
            },active: function() {
            },unactive: function() {
            }}, f);
        this._init();
    };

    b.prototype = {constructor: b,_init: function() {
            var g = this, f = g.options;
            this.zContainer = $(f.container).delegate("dd", "touchstart", function() {
                var i = $(this);
                if (i.hasClass("null") || i.hasClass("disable")) {
                    return;
                }
                var h = setTimeout(function() {
                    i.addClass("hover");
                }, 50);
                $(document).one("touchmove touchend", function() {
                    clearTimeout(h);
                    i.removeClass("hover");
                });
            }).delegate("dd", "click", function() {
                g.unactive(this);
            });
            f.init.call(this);
        },_render: function(g, k) {
            var i = this, f = i.options;
            i.zContainer.append(_.template($(f.tpl).html(), {value: g, humanize: k}));
        },render: function(k) {
            var j = this, h = j.options, g = -1, f;
            if (k) {
                j.zContainer.html("");
            }
            j._render(g, '<恢复默认>');
            for(g = 0; g < h.data.length; g++){
                j._render(g, h.data[g]);
            }
            return j;
        },active: function(day) {
            var f = this.options.active;
            f && f.call(this, day);
            return this;
        },unactive: function(value) {
            var f = this.options.unactive;
            f && f.call(this, value);
            return this;
        }
    };
    $._QunarWorkType = b;
})(Zepto);
(function (){
    $.App = function(){
        var loading = $('#mask'), types = ['急1', '急2', '付班', '夜班', '下夜班'], special = {}, pages, getTodayType = function(){
        	var now = new Date(),
        		monthFirstDay = new Date(now.getFullYear(), now.getMonth(), 1);
            return (Math.floor((monthFirstDay.getTime() + 28800000) / 86400000) - 2) % types.length;
        }, i = getTodayType(), calendar = $('#calendarPage').page({
            init : function() {
                var me = this, holiday = {
                    "2015-04-04":"4", 
                    "2015-04-05":"清明节", 
                    "2015-04-06":"6", 
                    "2015-05-01":"劳动节", 
                    "2015-05-02":"2", 
                    "2015-05-03":"3", 
                    "2015-06-20":"端午节", 
                    "2015-06-21":"21", 
                    "2015-06-22":"22", 
                    "2015-08-20":"七夕",
                    "2015-09-27":"中秋节",
                    "2015-10-01":"国庆节",
                    "2015-10-02":"2",
                    "2015-10-03":"3",
                    "2015-10-04":"4",
                    "2015-10-05":"5", 
                    "2015-10-06":"6", 
                    "2015-10-07":"7", 
                    "2015-11-11":"光棍节",
                    "2015-12-24":"平安夜",
                    "2015-12-25":"圣诞节",
                }, tiaoxiu = {
                    "2015-08-20":"2015-08-20",
                    "2015-10-10":"2015-10-10",
                    "2015-11-11":"2015-11-11",
                    "2015-12-24":"2015-12-24",
                    "2015-12-25":"2015-12-25"
                }, DAYTOMILS = 1000 * 60 * 60 * 24;
                var calendar = new $._QunarCalendar({
                    container : '.qn_calendar',
                    serverDate: '',
                    validDays : 365,
                    // init : function(){
                    //     this._offset = 3;
                    // },
                    fill : function(date){
                        var key = $.getDateString(date),
                                className = '', html = '<p>' + date.getDate() + '</p>',
                				today = new Date(), end = today,
                                startDate = this.attr('startDate'), backDate = this.attr('backDate'),
                                color = '', type = 0;

                        if(date.getDay() == 6 || date.getDay() == 0) className = 'weekend';
                        if(holiday[key]) {
                            className = 'holiday';
                            html = '<p>' + holiday[key] + '</p>';
                        }
                        if(tiaoxiu[key]) {
                            className += ' weekday'
                        }

                        if(date.getFullYear() == today.getFullYear() && date.getMonth() == today.getMonth() && date.getDate() == today.getDate()) {
                            className = '';
                            html = '<p>今天</p>';
                        }

                        if(this.attr('start') && date - startDate == 0) {
                            //className = 'active';
                            className = holiday[key] ? 'holiday active' : 'active';
                        }
                        className += end - date >= DAYTOMILS ? ' disable' : '';

                        if(this.attr('back')) {
                            if(date - backDate == 0) className = holiday[key] ? 'holiday active' : 'active';
                            end = startDate;
                            className += end - date > 0 ? ' disable' : '';
                        }

                        //html = html + '<p class="tag ' + (price.isLowPrice == 1 ? 'qn_red' : 'qn_grey') + '">&yen;' + price.price  + '</p>';

                        if(special[key] !== undefined){
                            i = special[key];
                            className += ' changed'
                        }

                        type = i % types.length;

                        if(className.indexOf('disable') === -1){
                            className += ' work_type_' + type;
                        }

                        html += '<p class="tag qn_grey"' + color + '>' + types[type] + '</p>';
                        i += 1;

                        return {
                            className : className,
                            html : html
                        }
                    },
                    active : function(date){
                        loading.show();
                        i = getTodayType();
                        $.ajax({
                            type: 'GET', 
                            url: 'plan', 
                            dataType: 'json',
                            cache: false,
                            success: function(data){
                                special = data;
                                calendar.render(null, null, { clear : true });
                                loading.hide();
                            },
                            error: function(xhr, type){
                                alert('请求出错!')
                                loading.hide();
                            }
                        });
                    },
                    unactive : function(dom){
                        var that = this, dom = $(dom), temp, day = dom.data('day');
                        if(dom.hasClass('disable') || dom.hasClass('null') || dom.attr('colspan')) return;
                        $('.qn_calendar').find('.active').removeClass('active');
                        //dom.addClass('active');
                        me.unactive();
                    }
                });
                this.calendar = calendar;
            },
            active: function(){
                this.calendar.active();
            }
        }),
        typePage = $('#typePage').page({
            init: function(){
                var me = this, changeWorkType = function(day, value){
                    loading.show();
                    $.ajax({
                        type: "POST", 
                        url: 'plan/set',
                        data: { day: day, value: value },
                        dataType: 'json',
                        cache: false,
                        success: function(data){
                            calendar.active();
                            loading.hide();
                        },
                        error: function(xhr, type){
                            alert('请求出错!')
                            loading.hide();
                        }
                    });
                }, workType = new $._QunarWorkType({
                    data: types,
                    container: '.qn_worktype',
                    active: function(day){
                        var title = $('#typePage .title');
                        workType.render(true)
                        title.html(day + '上班类型');
                        workType.day = day;
                        pages.toPage('typePage');
                    },
                    unactive: function(el){
                        var a = $(el), b = parseInt(a.data('value'));
                        changeWorkType(workType.day, b);
                        me.unactive();
                    }
                });

                // 点击日历日期弹出选择类型页面
                $('#calendarPage').delegate('td', 'click', function(){
                    var a = $(this), b = a.data('day');
                    if (a.hasClass("null") || a.hasClass("disable")) {
                        return;
                    }
                    workType.active(b);
                });

            }
            
        });

        calendar.init();
        calendar.active();
        var pages = $('.qn_pages').pages({
            main : 'calendarPage',
            side : [typePage]
        });
    };
})();