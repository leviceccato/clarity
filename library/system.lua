return function()
    local s = {}
    s.entities = {}
    s.components = {}

    s.addEntity = function(entity)
        s.entities[#s.entities + 1] = entity
    end

    s.sortEntities = function() end
    s.load = function() end
    s.update = function() end
    s.draw = function() end

    return s
end