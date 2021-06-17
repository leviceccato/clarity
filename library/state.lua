return function(arg)
    local s = {}
    s.worlds = {}
    s.activateWorld = nil
    s.controls = {}

    local keyMap = {
        ['space'] = 'jump',
        ['up'] = 'up',
        ['w'] = 'up',
        ['left'] = 'left',
        ['a'] = 'left',
        ['right'] = 'right',
        ['d'] = 'right',
        ['down'] = 'down',
        ['s'] = 'down',
        ['escape'] = 'menu',
        ['`'] = 'debug'
    }
    for _, control in pairs(keyMap) do
        s.controls[control] = false
    end

    s.updateControls = function(key, isDown)
        local control = keyMap[key]
        if control then
            s.controls[control] = isDown
        end
    end

    s.addWorld = function(world)
        s.worlds[world.name] = world
    end

    s.activateWorld = function(worldName)
        if s.activateWorld then
            s.activateWorld.exit()
        end
        s.worlds[worldName].load(arg)
        s.activateWorld = s.worlds[worldName]
    end

    s.update = function(dt)
        s.activateWorld.update(dt)
    end

    s.draw = function()
        s.activateWorld.draw()
    end

    s.keypressed = function(key)
        s.updateControls(key, true)
        s.activateWorld.keypressed(key)
    end

    s.keyreleased = function(key)
        s.updateControls(key, false)
        s.activateWorld.keyreleased(key)
    end

    s.resize = function(w, h)
        s.activateWorld.resize(w, h)
    end

    return s
end