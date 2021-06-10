local system = require('library.system')

return function()
    local s = system({ 'appearance', 'position', 'ui' })

    s.draw = function()
        love.graphics.print('FPS: ' .. love.timer.getFPS(), 10, 10)
    end

    return s
end