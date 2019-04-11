/* By Viktor A. Rozenko Voitenko (2019)
 *
 * This is the "objects.hpp" file. It contains all the main objects used by
 * the Scoops Virtual Machine (SVM) and probably various other modules.
 *
 * You must remember that the only data type that exists in the Scoops 
 * Programming Language is the "ScpObj" class and all other data types are
 * merely its subclasses that inherit its rather humble functionality and
 * expand on it. 
 *
 * Such architecture allows SVM to have a single data stack
 * that deals with all the values it wishes to deal with. If there were
 * multiple classes, it would be impossible to do that as C++ std::stack
 * can only have values of the same type pushed to it..
 *
 */



#pragma once

#include <string>



class ScpObj 
{ 
    std::string type;

    public:
    void set_type(std::string type) { this->type = type; }
    std::string get_type() { return this->type; }
};


class ScpBln : public ScpObj
{
    bool value;

    public:
    ScpBln (bool value)
    {
        this->set_type("bln");
        this->value = value;
    }
    
    bool get_value() { return this->value; }
};


class ScpChr : public ScpObj
{
    char value;

    public:
    ScpChr (char value)
    {
        this->set_type("chr");
        this->value = value;
    }

    char get_value() { return this->value; }
};


class ScpNum : public ScpObj
{
    double value;

    public:
    ScpNum (double value)
    {
        this->set_type("num");
        this->value = value;
    }

    double get_value() { return this->value; }
};


class ScpStr : public ScpObj
{
    std::string value;

    public:
    ScpStr (std::string value)
    {
        this->set_type("str");
        this->value = value;
    }

    std::string get_value() { return this->value; }

    void concat(ScpStr appendix) { this->value += appendix.get_value(); }
};

