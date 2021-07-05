local system = require('library.system')

return function(state)
    local s = system({'event'})

    local actions = {
        ['quit'] = function(data)
            local status = data.status or 0
            love.event.quit(status)
        end,
        ['activate-world'] = function(data)
            state.activateWorld(data.name)
        end
    }

    s.draw = function()
        local e
        -- Remove the last entity and run its event until none are left
        while #s.entities > 0 do
            e = table.remove(s.entities)
            actions[e.event.name](e.event.data)
        end
    end

    return s
end