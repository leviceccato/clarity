local system = require('library.system')

local timer = love.timer

return function(state)
    local s = system({'text', 'debug'})

    s.update = function(dt)
        for index = 1, #s.entities do
            local e = s.entities[index]
            e.text.content = 'FPS: ' .. timer.getFPS()
        end
    end

    return s
end