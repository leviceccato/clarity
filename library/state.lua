return function()
    local s = {}
    s.worlds = {}
    s.activateWorld = nil

    s.addWorld = function(world)
        s.worlds[world.name] = world
    end

    s.activateWorld = function(worldName)
        s.activateWorld = s.worlds[worldName]
    end

    s.load = function()
        s.activateWorld.load()
    end

    s.update = function()
        s.activateWorld.update()
    end

    s.draw = function()
        s.activateWorld.draw()
    end

    return s
end