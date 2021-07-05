local system = require('library.system')

return function(state)
    local s = system({'controls'})

    local updateInputs = function(inputName, inputData)
        local input = state.controls[inputName]
        if input then
            state.inputs[input] = inputData
        end
    end

    s.mousepressed = function(x, y, button, isTouch, pressCount)
        updateInputs('mouse' .. button, {
            x = x,
            y = y,
            isTouch = isTouch,
            pressCount = pressCount
        })
    end

    s.mousereleased = function(button)
        updateInputs('mouse' .. button, nil)
    end

    s.keypressed = function(key, isRepeat)
        updateInputs(key, {isRepeat = isRepeat})
    end

    s.keyreleased = function(key)
        updateInputs(key, nil)
    end

    return s
end