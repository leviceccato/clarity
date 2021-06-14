local system = require('library.system')

local floor = math.floor

return function(state)
    local s = system({'animation'})

    s.update = function(dt)
        for index = 1, #s.entities do
            local e = s.entities[index]
            e.animation.time = e.animation.time + (dt * 1000)
            if e.animation.time >= e.animation.duration then
                e.animation.time = e.animation.time - e.animation.duration
            end
            e.animation.frame = floor(e.animation.time / e.animation.duration * #e.animation.frames) + 1
        end
    end

    return s
end