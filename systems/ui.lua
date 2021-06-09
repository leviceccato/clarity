local system = require('library.system')

return function()
    local s = system()
    s.components = { 'appearance', 'position', 'ui' }

    s.draw = function()
        love.graphics.print('FPS: ' .. tostring(love.timer.getFPS()), 10, 10)
        -- local e
        -- for index = 1, #s.entities do
        --     e = s.entities[index]
        --     ...
        -- end
    end

    return s
end