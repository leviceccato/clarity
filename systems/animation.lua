local system = require('library.system')

return function()
    local s = system({'appearance', 'animation'})

    s.update = function(dt)
        for index = 1, #s.entities do
            local e = s.entities[index]
            e.appearance.time = e.appearance.time + dt
            if e.appearance.time >= e.appearance.duration then
                e.appearance.time = e.appearance.time - e.appearance.duration
            end
        end
    end

    return s
end