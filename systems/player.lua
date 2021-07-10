local system = require('library.system')

return function(state)
    local s = system({'position', 'controls'})

    s.update = function(dt)
        local distance = 200 * dt
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            if state.inputs.menu then
                state.activateWorlds({'title'})
            end
            if state.inputs.up then
                e.position.y = e.position.y - distance
            end
            if state.inputs.left then
                e.position.x = e.position.x - distance
            end
            if state.inputs.down then
                e.position.y = e.position.y + distance
            end
            if state.inputs.right then
                e.position.x = e.position.x + distance
            end
        end
    end

    return s
end