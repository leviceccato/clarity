return function(components)
    local s = {}
    s.entities = {}
    s.components = components or {}

    s.addEntity = function(entity)
        s.entities[#s.entities + 1] = entity
    end

    s.sortEntities = function() end

    s.load = function(arg) end
    s.update = function(dt) end
    s.draw = function() end
    s.keypressed = function(key) end
    s.keyreleased = function(key) end
    s.resize = function(width, height) end

    return s
end