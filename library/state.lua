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

    s.load = function(arg)
        s.activateWorld.load(arg)
    end

    s.update = function(dt)
        s.activateWorld.update(dt)
    end

    s.draw = function()
        s.activateWorld.draw()
    end

    s.keyreleased = function(...)
        s.activateWorld.keyreleased(...)
    end

    s.resize = function(...)
        s.activateWorld.resize(...)
    end

    return s
end