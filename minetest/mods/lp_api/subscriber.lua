--
-- subscriber
--
lp_api.subscriber = {}
lp_api.subscriber.category = minetest.settings:get("lp_api.subscriber.category") or "minetest"
lp_api.subscriber.timeout = minetest.settings:get("lp_api.subscriber.timeout") or 26

function lp_api.subscriber.sub_msg()
    lp_api.request_http("/sub?timeout="..lp_api.subscriber.timeout.."&category="..lp_api.subscriber.category, "",
    lp_api.subscriber.resp_handler)
end

function lp_api.subscriber.resp_handler(result)
    if result.succeeded and result.code == 200 then
        lp_api.subscriber.sub_msg()
        local data = minetest.parse_json(string.sub(result.data, string.find(result.data, "data")+6, -4))
        if data then
            lp_api.msg_router(data)
        end
    else
        minetest.after(lp_api.timeout, lp_api.subscriber.sub_msg)
        minetest.log("[lp_api] Sub request... [ERROR]")
    end
end

minetest.after(3, lp_api.subscriber.sub_msg)
