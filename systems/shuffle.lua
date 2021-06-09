local system = require('library.system')

return function()
    local s = system()
    s.components = { 'position', 'shuffle' }

    s.update = function()
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            e.position.x = e.position.x + (love.math.random() * 2 - 1)
            e.position.y = e.position.y + (love.math.random() * 2 - 1)
        end
    end

    return s
end