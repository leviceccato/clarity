return function()
    local s = {}
    s.entities = {}
    s.components = {}

    s.addEntity = function(entity)
        s.entities[#s.entities + 1] = entity
    end

    s.sortEntities = function() end

    s.load = function(arg) end
    s.update = function(dt) end
    s.draw = function() end
    s.keyreleased = function(...) end
    s.resize = function(...) end

    return s
end