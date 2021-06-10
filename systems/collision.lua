local system = require('library.system')

return function()
    local s = system({'appearance', 'position'})

    s.draw = function()
        local e
        for index = 1, #s.entities do
            e = s.entities[index]
            love.graphics.draw(e.appearance.image, e.position.x, e.position.y)
        end
    end

    return s
end