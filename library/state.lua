return function(arg)
    local s = {}
    s.worlds = {}
    s.activeWorld = nil

    s.inputs = {}
    s.controls = {
        ['mouse1'] = 'click',
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
    for _, input in pairs(s.controls) do
        s.inputs[input] = nil
    end

    s.addWorld = function(world)
        world.load(arg)
        s.worlds[world.name] = world
    end

    s.activateWorld = function(worldName)
        s.activeWorld = s.worlds[worldName]
    end

    s.update = function(dt)
        s.activeWorld.update(dt)
    end

    s.draw = function()
        s.activeWorld.draw()
    end

    s.exit = function()
        s.activeWorld.exit()
    end

    s.mousepressed = function(x, y, button, isTouch, pressCount)
        s.activeWorld.mousepressed(x, y, button, isTouch, pressCount)
    end

    s.mousereleased = function(button)
        s.activeWorld.mousereleased(button)
    end

    s.keypressed = function(key, isRepeat)
        s.activeWorld.keypressed(key, isRepeat)
    end

    s.keyreleased = function(key)
        s.activeWorld.keyreleased(key)
    end

    s.resize = function(w, h)
        s.activeWorld.resize(w, h)
    end

    return s
end