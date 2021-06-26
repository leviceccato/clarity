local system = require('library.system')

return function(state)
    local s = system({'controls'})

    local controls = {}
    local inputMap = {
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
        ['escape'] = 'escape',
        ['`'] = 'debug'
    }
    for _, control in pairs(inputMap) do
        controls[control] = nil
    end

    local updateControls = function(inputName, input)
        local control = inputMap[inputName]
        if control then
            controls[control] = input
        end
    end

    s.update = function(dt)
        for index = 1, #s.entities do
            local e = s.entities[index]
            if e.controls
            e.controls = controls
        end
    end

    s.mousepressed = function(x, y, button, isTouch, pressCount)
        updateControls('mouse' .. button, {
            x = x,
            y = y,
            isTouch = isTouch,
            pressCount = pressCount
        })
    end

    s.mousereleased = function(_, _, button)
        updateControls('mouse' .. button, nil)
    end

    s.keypressed = function(key, _, isRepeat)
        updateControls(key, {isRepeat = isRepeat})
    end

    s.keyreleased = function(key)
        updateControls(key, nil)
    end

    return s
end