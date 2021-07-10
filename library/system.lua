return function(components)
    local s = {}
    s.entities = {}
    s.components = components or {}

    s.addEntity = function(entity)
        s.entities[#s.entities + 1] = entity
    end

    s.sortEntities = function() end
    s.enter = function() end
    s.exit = function() end

    s.load = function(arg) end
    s.quit = function() end
    s.update = function(dt) end
    s.draw = function() end
    s.mousepressed = function(x, y, button, isTouch, pressCount) end
    s.mousereleased = function(button) end
    s.keypressed = function(key, isRepeat) end
    s.keyreleased = function(key) end
    s.resize = function(width, height) end

    return s
end